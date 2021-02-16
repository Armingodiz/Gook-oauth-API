package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArminGodiz/Gook-oauth-API/src/clients/redis_db"
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/access_token"
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetById(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	at, err := redis_db.DB.HGet(context.Background(), "tokens", accessTokenID).Result()
	if err != nil {
		return nil, errors.NewInternalServerError("no access token found !")
	}
	fmt.Println(at)
	var accessToken access_token.AccessToken
	json.Unmarshal([]byte(at), &accessToken)
	return &accessToken, nil
}
func (db *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	json1, err := json.Marshal(at)
	if err != nil {
		return errors.NewBadRequestError("invalid form")
	}
	err = redis_db.DB.HSet(context.Background(), "tokens", at.AccessToken, json1).Err()
	if err != nil {
		return errors.NewInternalServerError("error while saving in db")
	}
	return nil
}

func (db *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	_ = redis_db.DB.HDel(context.Background(), "tokens", at.AccessToken)
	db.Create(at)
	return nil
}
