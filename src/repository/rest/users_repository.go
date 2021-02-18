package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/users"
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
	"github.com/go-resty/resty/v2"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}
type restUsersRepository struct {
}

var (
	restyClient = resty.New()
)

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

func (ur *restUsersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {
	//response := UsersRestClient.Post("/users/login", users.LoginRequest{Email: email, Password: password})
	resp, err := restyClient.R().
		SetBody(users.LoginRequest{Email: email, Password: password}).
		SetResult(&users.User{}). // or SetResult(AuthSuccess{}).
		SetError(&errors.RestErr{}). // or SetError(AuthError{}).
		Post("http://localhost:1111/users/login")
	if err != nil {
		return nil, errors.NewInternalServerError("error while sending post request for login :     " + err.Error())
	}
	if resp == nil || resp.RawResponse == nil {
		return nil, errors.NewInternalServerError("INVALID email or password for login !")
	}
	if resp.StatusCode() > 299 { // we have an error
		var restErr errors.RestErr
		err := json.Unmarshal(resp.Body(), &restErr)
		fmt.Println(resp.String())
		if err != nil { // we get a different type of error
			return nil, errors.NewInternalServerError("Unknown error type accrued while trying to login ==>" + err.Error())
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error in unmarshalling response ")
	}
	return &user, nil
}
