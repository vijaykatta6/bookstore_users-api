package users_domain

import (
	"fmt"

	rest_errors "github.com/vijaykatta6/bookstore-users-api/utils/errors"
)

// access layer to out database

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *rest_errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return rest_errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *rest_errors.RestErr {

	// checking that user already exists or not, if exists then throw an bad request
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return rest_errors.NewBadRequestError(fmt.Sprintf("user %s already registered", user.Email))
		}
		return rest_errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	userDB[user.Id] = user
	return nil

}
