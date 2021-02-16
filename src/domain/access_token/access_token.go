package access_token

import (
	"fmt"
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token")
	} else if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	} else if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	} else if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expires")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println(expirationTime)
	return expirationTime.Before(now)
}
