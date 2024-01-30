package users_domain

import (
	"strings"

	rest_errors "github.com/vijaykatta6/bookstore-users-api/utils/errors"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

func (user *User) Validate() *rest_errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}
	return nil

}
