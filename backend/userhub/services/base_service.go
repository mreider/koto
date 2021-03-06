package services

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"path/filepath"
	"strings"

	"github.com/ansel1/merry"
	"github.com/disintegration/imaging"

	"github.com/mreider/koto/backend/common"
	"github.com/mreider/koto/backend/token"
	"github.com/mreider/koto/backend/userhub/caches"
	"github.com/mreider/koto/backend/userhub/config"
	"github.com/mreider/koto/backend/userhub/repo"

	"github.com/h2non/filetype"
)

type BaseService struct {
	repos              repo.Repos
	userCache          caches.Users
	s3Storage          *common.S3Storage
	tokenGenerator     token.Generator
	tokenParser        token.Parser
	mailSender         *common.MailSender
	cfg                config.Config
	notificationSender NotificationSender
	rootEmailTemplate  *template.Template
}

type BaseServiceOptions struct {
	UserCache          caches.Users
	S3Storage          *common.S3Storage
	TokenGenerator     token.Generator
	TokenParser        token.Parser
	MailSender         *common.MailSender
	Cfg                config.Config
	NotificationSender NotificationSender
	RootEmailTemplate  *template.Template
}

func NewBase(repos repo.Repos, options BaseServiceOptions) *BaseService {
	return &BaseService{
		repos:              repos,
		userCache:          options.UserCache,
		s3Storage:          options.S3Storage,
		tokenGenerator:     options.TokenGenerator,
		tokenParser:        options.TokenParser,
		mailSender:         options.MailSender,
		cfg:                options.Cfg,
		notificationSender: options.NotificationSender,
		rootEmailTemplate:  options.RootEmailTemplate,
	}
}

func (s *BaseService) getMe(ctx context.Context) repo.User {
	return ctx.Value(ContextUserKey).(repo.User)
}

func (s *BaseService) hasMe(ctx context.Context) bool {
	_, ok := ctx.Value(ContextUserKey).(repo.User)
	return ok
}

func (s *BaseService) isAdmin(ctx context.Context) bool {
	isAdmin, _ := ctx.Value(ContextIsAdminKey).(bool)
	return isAdmin
}

func (s *BaseService) getGroup(ctx context.Context, groupID string) (*repo.Group, bool) {
	group := s.repos.Group.FindGroupByID(groupID)
	if group == nil {
		return nil, false
	}
	me := s.getMe(ctx)
	return group, group.AdminID == me.ID
}

func (s *BaseService) getUser(ctx context.Context, userID string) (*repo.User, bool) {
	user := s.repos.User.FindUserByID(userID)
	if user == nil {
		return nil, false
	}
	me := s.getMe(ctx)
	isFriend := s.repos.Friend.AreFriends(me.ID, user.ID)
	return user, isFriend
}

func (s *BaseService) GetUserAttachments(ctx context.Context, userID string) common.MailAttachmentList {
	me := s.getMe(ctx)
	userInfo := s.userCache.User(userID, me.ID)
	if userInfo.AvatarThumbnailID == "" {
		return nil
	}

	var b bytes.Buffer
	err := s.s3Storage.Read(ctx, userInfo.AvatarThumbnailID, &b)
	if err != nil {
		log.Println("can't read user avatar:", err)
		return nil
	}

	mimeType := ""
	fileName := "avatar"
	t, _ := filetype.Match(b.Bytes())
	if t != filetype.Unknown {
		mimeType = t.MIME.Value
		if t.Extension != "" {
			fileName += "." + t.Extension
		}
	}

	return common.MailAttachmentList{
		"avatar": {
			Inline:   true,
			Data:     b.Bytes(),
			FileName: fileName,
			MIMEType: mimeType,
		},
	}
}

func (s *BaseService) saveThumbnail(ctx context.Context, avatarID string) (string, error) {
	var buf bytes.Buffer
	err := s.s3Storage.Read(ctx, avatarID, &buf)
	if err != nil {
		return "", merry.Wrap(err)
	}
	dataType, err := filetype.Match(buf.Bytes())
	if err != nil {
		return "", merry.Wrap(err)
	}
	if dataType.MIME.Type != "image" {
		return "", merry.New("not image")
	}

	orientation := common.GetImageOrientation(bytes.NewReader(buf.Bytes()))
	original, err := common.DecodeImageAndFixOrientation(bytes.NewReader(buf.Bytes()), orientation)
	if err != nil {
		return "", merry.Wrap(err)
	}
	thumbnail := imaging.Thumbnail(original, avatarThumbnailWidth, avatarThumbnailHeight, imaging.Lanczos)
	buf.Reset()
	err = imaging.Encode(&buf, thumbnail, imaging.JPEG)
	if err != nil {
		return "", merry.Wrap(err)
	}

	ext := filepath.Ext(avatarID)
	thumbnailID := strings.TrimSuffix(avatarID, ext) + "-thumbnail.jpg"
	err = s.s3Storage.PutObject(ctx, thumbnailID, buf.Bytes(), "image/jpeg")
	if err != nil {
		return "", merry.Wrap(err)
	}
	return thumbnailID, nil
}
