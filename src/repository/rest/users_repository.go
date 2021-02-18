package rest

import (
	"encoding/json"
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/users"
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
	"github.com/golang-restclient/rest"
	"time"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}
type restUsersRepository struct {
}

var (
	UsersRestClient = rest.RequestBuilder{
		BaseURL: "localhost:1111", // ** the port which users api is listening on **
		Timeout: 2 * time.Second,
	}
)

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

func (ur *restUsersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {
	response := UsersRestClient.Post("/users/login", users.LoginRequest{Email: email, Password: password})
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("INVALID email or password for login !")
	}
	if response.StatusCode > 299 { // we have an error
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil { // we get a different type of error
			return nil, errors.NewInternalServerError("Unknown error type accrued while trying to login ")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error in unmarshalling response ")
	}
	return &user, nil
}
