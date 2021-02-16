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
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	at, err := redis_db.DB.HGet(context.Background(), "tokens", id).Result()
	if err != nil {
		//fmt.Println(err)
		return nil, errors.NewInternalServerError("Error in db")
	}
	fmt.Println(at)
	var accessToken access_token.AccessToken
	json.Unmarshal([]byte(at), &accessToken)
	return &accessToken, nil
}
