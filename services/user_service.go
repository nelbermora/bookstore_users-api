package services

import (
	"github.com/nelbermora/bookstore_users-api/domain/users"
	"github.com/nelbermora/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	err := user.Validate()
	if err != nil {
		return nil, err
	}
	err = user.Save()
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func GetUser(userId int64) (*users.User, *errors.RestErr) {

	result := users.User{Id: userId}
	err := result.Get()
	if err != nil {
		return nil, err
	}
	return &result, nil

}
