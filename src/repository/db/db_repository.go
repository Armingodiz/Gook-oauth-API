package db

import (
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/access_token"
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("data base connection is not implemented !")
}
