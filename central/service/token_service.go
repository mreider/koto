package service

import (
	"time"

	"github.com/mreider/koto/central/repo"
	"github.com/mreider/koto/token"
)

type TokenService interface {
	Auth(user repo.User) (token string, err error)
	PostMessage(user repo.User, communities []string) (tokens []string, err error)
	GetMessages(user repo.User, communities []string) (tokens []string, err error)
}

type tokenService struct {
	relationRepo   repo.RelationRepo
	tokenGenerator token.Generator
}

func NewToken(relationRepo repo.RelationRepo, tokenGenerator token.Generator) TokenService {
	return &tokenService{
		relationRepo:   relationRepo,
		tokenGenerator: tokenGenerator,
	}
}

func (s *tokenService) Auth(user repo.User) (token string, err error) {
	return s.tokenGenerator.Generate(user, "auth", time.Now().Add(time.Minute*10), nil)
}

func (s *tokenService) PostMessage(user repo.User, communities []string) (tokens []string, err error) {
	userCommunities, err := s.relationRepo.InvitedCommunities(user)
	if err != nil {
		return nil, err
	}
	communitySet := make(map[string]bool)
	for _, community := range userCommunities {
		communitySet[community] = true
	}

	tokens = make([]string, len(communities))
	exp := time.Now().Add(time.Minute * 10)
	for i, community := range communities {
		if communitySet[community] {
			claims := map[string]interface{}{"community": community}
			communityToken, err := s.tokenGenerator.Generate(user, "post-message", exp, claims)
			if err != nil {
				return nil, err
			}
			tokens[i] = communityToken
		} else {
			tokens[i] = ""
		}
	}
	return tokens, nil
}

func (s *tokenService) GetMessages(user repo.User, communities []string) (tokens []string, err error) {
	relatedCommunities, err := s.relationRepo.RelatedCommunities(user)
	if err != nil {
		return nil, err
	}
	communityMap := make(map[string]repo.Community)
	for _, community := range relatedCommunities {
		communityMap[community.Address] = community
	}

	tokens = make([]string, len(communities))
	exp := time.Now().Add(time.Minute * 10)
	for i, community := range communities {
		if relatedCommunity, ok := communityMap[community]; ok {
			claims := map[string]interface{}{
				"community": relatedCommunity.Address,
				"users":     relatedCommunity.Friends,
			}
			communityToken, err := s.tokenGenerator.Generate(user, "get-messages", exp, claims)
			if err != nil {
				return nil, err
			}
			tokens[i] = communityToken
		} else {
			tokens[i] = ""
		}
	}
	return tokens, nil
}