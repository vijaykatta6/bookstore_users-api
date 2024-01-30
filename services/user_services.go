package services

import (
	users_domain "github.com/vijaykatta6/bookstore-users-api/Domain/users"
	rest_errors "github.com/vijaykatta6/bookstore-users-api/utils/errors"
)

func GetUSer(userId int64) (*users_domain.User, *rest_errors.RestErr) {

	result := &users_domain.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user users_domain.User) (*users_domain.User, *rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
