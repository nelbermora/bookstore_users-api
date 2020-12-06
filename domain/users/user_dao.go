package users

import (
	"fmt"

	"github.com/nelbermora/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("User %d nor found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.Lastanme = result.Lastanme
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		return errors.NewBadRequestErr(fmt.Sprintf("User %d already exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil

}
