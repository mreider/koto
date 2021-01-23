package services

import (
	"context"
	"regexp"
	"strings"

	"github.com/ansel1/merry"
	"github.com/twitchtv/twirp"

	"github.com/mreider/koto/backend/common"
	"github.com/mreider/koto/backend/userhub/repo"
	"github.com/mreider/koto/backend/userhub/rpc"
)

var (
	groupNameRe = regexp.MustCompile(`^(\p{L}|\d)(\p{L}|\d|-|_|\.| )+(\p{L}|\d)$`)
)

func NewGroup(base *BaseService) rpc.GroupService {
	return &groupService{
		BaseService: base,
	}
}

type groupService struct {
	*BaseService
}

func (s *groupService) ManagedGroup(ctx context.Context, _ *rpc.Empty) (*rpc.GroupManagedGroupsResponse, error) {
	me := s.getMe(ctx)
	groups := s.repos.Group.ManagedGroups(me.ID)
	rpcGroups := make([]*rpc.Group, len(groups))
	for i, group := range groups {
		rpcGroups[i] = &rpc.Group{
			Id:          group.ID,
			Name:        group.Name,
			Description: group.Description,
			IsPublic:    group.IsPublic,
			MemberCount: int32(group.MemberCount),
		}
	}

	return &rpc.GroupManagedGroupsResponse{
		Groups: rpcGroups,
	}, nil
}

func (s *groupService) AddGroup(ctx context.Context, r *rpc.GroupAddGroupRequest) (*rpc.GroupAddGroupResponse, error) {
	me := s.getMe(ctx)

	if r.Name == "" {
		return nil, twirp.InvalidArgumentError("name", "shouldn't be empty")
	}
	if r.Description == "" {
		return nil, twirp.InvalidArgumentError("description", "shouldn't be empty")
	}
	r.Name = strings.TrimSpace(r.Name)
	if !groupNameRe.MatchString(r.Name) {
		return nil, twirp.InvalidArgumentError("name", "is invalid")
	}

	group := s.repos.Group.FindGroupByName(r.Name)
	if group != nil {
		return nil, twirp.NewError(twirp.AlreadyExists, "group already exists")
	}

	groupID := common.GenerateUUID()

	s.repos.Group.AddGroup(groupID, r.Name, r.Description, me.ID, r.IsPublic)

	group = s.repos.Group.FindGroupByID(groupID)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}

	if r.AvatarId != "" {
		err := s.setAvatar(ctx, *group, r.AvatarId)
		if err != nil {
			return nil, err
		}
	}

	if r.BackgroundId != "" {
		s.repos.Group.SetBackground(group.ID, r.BackgroundId)
	}

	backgroundLink, err := s.createBlobLink(ctx, r.BackgroundId)
	if err != nil {
		return nil, err
	}

	return &rpc.GroupAddGroupResponse{
		Group: &rpc.Group{
			Id:          group.ID,
			Name:        group.Name,
			Description: group.Description,
			IsPublic:    group.IsPublic,
			Background:  backgroundLink,
		},
	}, nil
}

func (s *groupService) EditGroup(ctx context.Context, r *rpc.GroupEditGroupRequest) (*rpc.Empty, error) {
	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}
	if !isGroupAdmin {
		return nil, twirp.NewError(twirp.PermissionDenied, "")
	}

	if r.DescriptionChanged {
		if r.Description == "" {
			return nil, twirp.InvalidArgumentError("description", "is empty")
		}
		if r.Description != group.Description {
			s.repos.Group.SetDescription(group.ID, r.Description)
		}
	}

	if r.IsPublicChanged {
		s.repos.Group.SetIsPublic(group.ID, r.IsPublic)
	}

	if r.AvatarChanged {
		err := s.setAvatar(ctx, *group, r.AvatarId)
		if err != nil {
			return nil, merry.Wrap(err)
		}
	}

	if r.BackgroundChanged {
		s.repos.Group.SetBackground(group.ID, r.BackgroundId)
	}

	return &rpc.Empty{}, nil
}

func (s *groupService) DeleteGroup(ctx context.Context, r *rpc.GroupDeleteGroupRequest) (*rpc.Empty, error) {
	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}
	if !isGroupAdmin {
		return nil, twirp.NewError(twirp.PermissionDenied, "")
	}

	s.repos.Group.DeleteGroup(r.GroupId)
	return &rpc.Empty{}, nil
}

func (s *groupService) setAvatar(ctx context.Context, group repo.Group, avatarID string) (err error) {
	if group.AvatarOriginalID == avatarID {
		return nil
	}

	if avatarID == "" {
		s.repos.Group.SetAvatar(group.ID, "", "")
		return nil
	}

	thumbnailID, err := s.saveThumbnail(ctx, avatarID)
	if err != nil {
		return merry.Wrap(err)
	}

	s.repos.Group.SetAvatar(group.ID, avatarID, thumbnailID)
	return nil
}

func (s *groupService) AddUser(ctx context.Context, r *rpc.GroupAddUserRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	if me.ID == r.UserId {
		return nil, twirp.InvalidArgumentError("user_id", "")
	}

	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil || !isGroupAdmin {
		return nil, twirp.NotFoundError("group not found")
	}

	if s.repos.Group.IsGroupMember(r.GroupId, r.UserId) {
		return nil, twirp.NewError(twirp.AlreadyExists, "already in the group.")
	}

	if s.repos.Friend.AreFriends(me.ID, r.UserId) {
		s.repos.Group.AddUserToGroup(r.GroupId, r.UserId)
	} else {
		_, err := s.CreateInvite(ctx, &rpc.GroupCreateInviteRequest{
			GroupId: r.GroupId,
			Invited: r.UserId,
		})
		if err != nil {
			return nil, err
		}
	}

	return &rpc.Empty{}, nil
}

func (s *groupService) RequestJoin(ctx context.Context, r *rpc.GroupRequestJoinRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	return s.CreateInvite(ctx, &rpc.GroupCreateInviteRequest{
		GroupId: r.GroupId,
		Invited: me.ID,
		Message: r.Message,
	})
}

func (s *groupService) CreateInvite(ctx context.Context, r *rpc.GroupCreateInviteRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)

	if r.GroupId == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "group_id")
	}
	if r.Invited == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "invited")
	}

	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil || (!group.IsPublic && !isGroupAdmin) {
		return nil, twirp.NotFoundError("group not found")
	}

	if me.ID != r.Invited && !s.repos.Group.IsGroupMember(group.ID, me.ID) {
		return nil, twirp.NewError(twirp.PermissionDenied, "not a group member")
	}

	invitedUser := s.repos.User.FindUserByIDOrName(r.Invited)
	if invitedUser == nil {
		if strings.Contains(r.Invited, "@") && strings.Contains(r.Invited, ".") {
			friends := s.repos.User.FindUsersByEmail(r.Invited)
			if len(friends) == 1 {
				invitedUser = &friends[0]
			} else if len(friends) > 1 {
				return nil, twirp.NewError(twirp.AlreadyExists, "Email belongs to more than one account. Invite by username instead.")
			}
		} else {
			return nil, twirp.NotFoundError("user not found")
		}
	}

	if invitedUser != nil {
		areBlocked := s.repos.User.AreBlocked(me.ID, invitedUser.ID)
		if areBlocked {
			return nil, twirp.NotFoundError("user not found")
		}

		if s.repos.Group.IsGroupMember(group.ID, invitedUser.ID) {
			return nil, twirp.NewError(twirp.AlreadyExists, "already in the group")
		}

		s.repos.Group.AddInvite(group.ID, me.ID, invitedUser.ID, r.Message)
		if s.notificationSender != nil {
			meInfo := s.userCache.UserFullAccess(me.ID)
			s.notificationSender.SendNotification([]string{invitedUser.ID}, meInfo.DisplayName+" invited you to the group "+group.Name+"", "group-invite/add", map[string]interface{}{
				"group_id": group.ID,
				"user_id":  me.ID,
			})
		}
		// TODO:
		//err = s.sendInviteLinkToRegisteredUser(ctx, user, invitedUser.Email)
		//if err != nil {
		//	log.Println("can't invite by email:", err)
		//}
	} else {
		s.repos.Group.AddInviteByEmail(group.ID, me.ID, r.Invited, r.Message)

		// TODO
		//err = s.sendInviteLinkToUnregisteredUser(ctx, user, r.Invited)
		//if err != nil {
		//	log.Println("can't invite by email:", err)
		//}
	}

	return &rpc.Empty{}, nil
}

func (s *groupService) AcceptInvite(ctx context.Context, r *rpc.GroupAcceptInviteRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)

	if r.GroupId == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "group_id")
	}

	group, _ := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}

	if !s.repos.Group.AcceptInvite(r.GroupId, r.InviterId, me.ID) {
		return nil, twirp.NotFoundError("invite not found")
	}

	if r.InviterId == group.AdminID {
		s.repos.Group.ConfirmInvite(r.GroupId, r.InviterId, me.ID)
	}

	return &rpc.Empty{}, nil
}

func (s *groupService) RejectInvite(ctx context.Context, r *rpc.GroupRejectInviteRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)

	if r.GroupId == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "group_id")
	}

	if !s.repos.Group.RejectInvite(r.GroupId, r.InviterId, me.ID) {
		return nil, twirp.NotFoundError("invite not found")
	}
	return &rpc.Empty{}, nil
}

func (s *groupService) InvitesFromMe(ctx context.Context, _ *rpc.Empty) (*rpc.GroupInvitesFromMeResponse, error) {
	me := s.getMe(ctx)
	invites := s.repos.Group.InvitesFromMe(me)
	rpcInvites := make([]*rpc.GroupInvite, len(invites))
	for i, invite := range invites {
		invitedInfo := s.userCache.User(invite.InvitedID, me.ID)
		invitedName, invitedFullName := invitedInfo.Name, invitedInfo.FullName
		if invite.InvitedID == "" {
			invitedName = invite.InvitedEmail
			invitedFullName = ""
		}

		rpcInvites[i] = &rpc.GroupInvite{
			GroupId:           invite.GroupID,
			GroupName:         invite.GroupName,
			GroupDescription:  invite.GroupDescription,
			InvitedId:         invite.InvitedID,
			InvitedName:       invitedName,
			InvitedFullName:   invitedFullName,
			CreatedAt:         common.TimeToRPCString(invite.CreatedAt),
			AcceptedAt:        common.NullTimeToRPCString(invite.AcceptedAt),
			RejectedAt:        common.NullTimeToRPCString(invite.RejectedAt),
			AcceptedByAdminAt: common.NullTimeToRPCString(invite.AcceptedByAdminAt),
			RejectedByAdminAt: common.NullTimeToRPCString(invite.RejectedByAdminAt),
			Message:           invite.Message,
		}
	}

	return &rpc.GroupInvitesFromMeResponse{
		Invites: rpcInvites,
	}, nil
}

func (s *groupService) InvitesForMe(ctx context.Context, _ *rpc.Empty) (*rpc.GroupInvitesForMeResponse, error) {
	me := s.getMe(ctx)
	invites := s.repos.Group.InvitesForMe(me)
	rpcInvites := make([]*rpc.GroupInvite, len(invites))
	for i, invite := range invites {
		inviterInfo := s.userCache.UserFullAccess(invite.InviterID)
		rpcInvites[i] = &rpc.GroupInvite{
			GroupId:           invite.GroupID,
			GroupName:         invite.GroupName,
			GroupDescription:  invite.GroupDescription,
			InviterId:         invite.InviterID,
			InviterName:       inviterInfo.Name,
			InviterFullName:   inviterInfo.FullName,
			CreatedAt:         common.TimeToRPCString(invite.CreatedAt),
			AcceptedAt:        common.NullTimeToRPCString(invite.AcceptedAt),
			RejectedAt:        common.NullTimeToRPCString(invite.RejectedAt),
			AcceptedByAdminAt: common.NullTimeToRPCString(invite.AcceptedByAdminAt),
			RejectedByAdminAt: common.NullTimeToRPCString(invite.RejectedByAdminAt),
			Message:           invite.Message,
		}
	}

	return &rpc.GroupInvitesForMeResponse{
		Invites: rpcInvites,
	}, nil
}

func (s *groupService) LeaveGroup(ctx context.Context, r *rpc.GroupLeaveGroupRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)

	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}
	if isGroupAdmin {
		return nil, twirp.NewError(twirp.InvalidArgument, "admin can't leave the group")
	}

	s.repos.Group.RemoveUserFromGroup(r.GroupId, me.ID)
	return &rpc.Empty{}, nil
}

func (s *groupService) RemoveUser(ctx context.Context, r *rpc.GroupRemoveUserRequest) (*rpc.Empty, error) {
	me := s.getMe(ctx)
	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}
	if !isGroupAdmin {
		return nil, twirp.NewError(twirp.PermissionDenied, "")
	}
	if me.ID == r.UserId {
		return nil, twirp.NewError(twirp.InvalidArgument, "can't delete himself")
	}

	s.repos.Group.RemoveUserFromGroup(r.GroupId, r.UserId)
	return &rpc.Empty{}, nil
}

func (s *groupService) ConfirmInvite(ctx context.Context, r *rpc.GroupConfirmInviteRequest) (*rpc.Empty, error) {
	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}
	if !isGroupAdmin {
		return nil, twirp.NewError(twirp.PermissionDenied, "")
	}

	s.repos.Group.ConfirmInvite(r.GroupId, r.InviterId, r.InvitedId)
	return &rpc.Empty{}, nil
}

func (s *groupService) DenyInvite(ctx context.Context, r *rpc.GroupDenyInviteRequest) (*rpc.Empty, error) {
	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}
	if !isGroupAdmin {
		return nil, twirp.NewError(twirp.PermissionDenied, "")
	}

	s.repos.Group.DenyInvite(r.GroupId, r.InviterId, r.InvitedId)
	return &rpc.Empty{}, nil
}

func (s *groupService) InvitesToConfirm(ctx context.Context, _ *rpc.Empty) (*rpc.GroupInvitesToConfirmResponse, error) {
	me := s.getMe(ctx)
	invites := s.repos.Group.AdminInvitesToConfirm(me.ID)

	groupSet := make(map[string]*rpc.GroupInvites)
	var groups []*rpc.GroupInvites
	for _, invite := range invites {
		inviterInfo := s.userCache.User(invite.InviterID, me.ID)
		invitedInfo := s.userCache.User(invite.InvitedID, me.ID)
		rpcInvite := &rpc.GroupInvite{
			InviterId:         invite.InviterID,
			InviterName:       inviterInfo.Name,
			InviterFullName:   inviterInfo.FullName,
			InvitedId:         invite.InvitedID,
			InvitedName:       invitedInfo.Name,
			InvitedFullName:   invitedInfo.FullName,
			CreatedAt:         common.TimeToRPCString(invite.CreatedAt),
			AcceptedAt:        common.NullTimeToRPCString(invite.AcceptedAt),
			RejectedAt:        common.NullTimeToRPCString(invite.RejectedAt),
			AcceptedByAdminAt: common.NullTimeToRPCString(invite.AcceptedByAdminAt),
			RejectedByAdminAt: common.NullTimeToRPCString(invite.RejectedByAdminAt),
			Message:           invite.Message,
		}

		if _, ok := groupSet[invite.GroupID]; !ok {
			gi := &rpc.GroupInvites{
				Group: &rpc.Group{
					Id:          invite.GroupID,
					Name:        invite.GroupName,
					Description: invite.GroupDescription,
					IsPublic:    invite.GroupIsPublic,
				},
				Invites: []*rpc.GroupInvite{rpcInvite},
			}
			groupSet[invite.GroupID] = gi
			groups = append(groups, gi)
		} else {
			groupSet[invite.GroupID].Invites = append(groupSet[invite.GroupID].Invites, rpcInvite)
		}
	}

	return &rpc.GroupInvitesToConfirmResponse{
		Groups: groups,
	}, nil
}

func (s *groupService) PublicGroups(ctx context.Context, _ *rpc.Empty) (*rpc.GroupPublicGroupsResponse, error) {
	me := s.getMe(ctx)

	groups := s.repos.Group.PublicGroups()
	statuses := s.repos.Group.JoinStatuses(me.ID)

	rpcGroups := make([]*rpc.GroupPublicGroupsResponseItem, len(groups))
	for i, group := range groups {
		adminInfo := s.userCache.User(group.AdminID, me.ID)
		rpcGroups[i] = &rpc.GroupPublicGroupsResponseItem{
			Group: &rpc.Group{
				Id:          group.ID,
				Name:        group.Name,
				Description: group.Description,
				IsPublic:    group.IsPublic,
				MemberCount: int32(group.MemberCount),
				Admin: &rpc.User{
					Id:           group.AdminID,
					Name:         adminInfo.Name,
					FullName:     adminInfo.FullName,
					HideIdentity: adminInfo.HideIdentity,
				},
			},
			Status: statuses[group.ID],
		}
	}

	return &rpc.GroupPublicGroupsResponse{
		Groups: rpcGroups,
	}, nil
}

func (s *groupService) GroupDetails(ctx context.Context, r *rpc.GroupGroupDetailsRequest) (*rpc.GroupGroupDetailsResponse, error) {
	me := s.getMe(ctx)
	group, isGroupAdmin := s.getGroup(ctx, r.GroupId)
	if group == nil {
		return nil, twirp.NotFoundError("group not found")
	}
	isMember := s.repos.Group.IsGroupMember(r.GroupId, me.ID)
	if !group.IsPublic && !isMember {
		return nil, twirp.NewError(twirp.PermissionDenied, "")
	}
	backgroundLink, err := s.createBlobLink(ctx, group.BackgroundID)
	if err != nil {
		return nil, err
	}
	adminInfo := s.userCache.User(group.AdminID, me.ID)
	rpcGroup := &rpc.Group{
		Id:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		IsPublic:    group.IsPublic,
		Background:  backgroundLink,
		MemberCount: int32(group.MemberCount),
		Admin: &rpc.User{
			Id:           group.AdminID,
			Name:         adminInfo.Name,
			FullName:     adminInfo.FullName,
			HideIdentity: adminInfo.HideIdentity,
		},
	}

	if !isMember {
		return &rpc.GroupGroupDetailsResponse{
			Group: rpcGroup,
		}, nil
	}

	members := s.repos.Group.GroupMembers(r.GroupId)
	rpcMembers := make([]*rpc.User, 0, len(members))
	for _, member := range members {
		memberInfo := s.userCache.User(member.ID, me.ID)
		rpcMembers = append(rpcMembers, &rpc.User{
			Id:           member.ID,
			Name:         memberInfo.Name,
			FullName:     memberInfo.FullName,
			HideIdentity: memberInfo.HideIdentity,
		})
	}

	var rpcInvites []*rpc.GroupInvite
	if isGroupAdmin {
		invites := s.repos.Group.GroupInvitesToConfirm(r.GroupId)

		for _, invite := range invites {
			inviterInfo := s.userCache.User(invite.InviterID, me.ID)
			invitedInfo := s.userCache.User(invite.InvitedID, me.ID)
			rpcInvite := &rpc.GroupInvite{
				InviterId:         invite.InviterID,
				InviterName:       inviterInfo.Name,
				InviterFullName:   inviterInfo.FullName,
				InvitedId:         invite.InvitedID,
				InvitedName:       invitedInfo.Name,
				InvitedFullName:   invitedInfo.FullName,
				CreatedAt:         common.TimeToRPCString(invite.CreatedAt),
				AcceptedAt:        common.NullTimeToRPCString(invite.AcceptedAt),
				RejectedAt:        common.NullTimeToRPCString(invite.RejectedAt),
				AcceptedByAdminAt: common.NullTimeToRPCString(invite.AcceptedByAdminAt),
				RejectedByAdminAt: common.NullTimeToRPCString(invite.RejectedByAdminAt),
				Message:           invite.Message,
			}
			rpcInvites = append(rpcInvites, rpcInvite)
		}
	}

	return &rpc.GroupGroupDetailsResponse{
		Group:   rpcGroup,
		Members: rpcMembers,
		Invites: rpcInvites,
	}, nil
}
