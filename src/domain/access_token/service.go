package access_token

import (
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("INVALID Access Token ID !")
	}
	accessToken, err := s.repository.GetById(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}


func (s *service) Create(at AccessToken) *errors.RestErr {
	err := at.Validate()
	if err != nil {
		return err
	}
	return s.repository.Create(at)
}
func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	err := at.Validate()
	if err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}