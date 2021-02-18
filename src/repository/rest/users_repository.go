package rest

import (
	"github.com/ArminGodiz/Gook-oauth-API/src/domain/users"
	"github.com/ArminGodiz/Gook-oauth-API/src/utils/errors"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}
type restUsersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

func (ur *restUsersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {

}
