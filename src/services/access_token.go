package services

import (
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/access_token"
	"github.com/ArminGodiz/Gook-oauth-API/src/repository/db"
	"github.com/ArminGodiz/Gook-oauth-API/src/repository/rest"
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}
type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}
type service struct {
	DBRepository   db.DbRepository
	RestRepository rest.RestUsersRepository
}

func NewService(db db.DbRepository, restRepo rest.RestUsersRepository) Service {
	return &service{
		DBRepository:   db,
		RestRepository: restRepo,
	}
}

func (s *service) GetById(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("INVALID Access Token ID !")
	}
	accessToken, err := s.DBRepository.GetById(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(req access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.RestRepository.LoginUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.DBRepository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}
func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	err := at.Validate()
	if err != nil {
		return err
	}
	return s.DBRepository.UpdateExpirationTime(at)
}
