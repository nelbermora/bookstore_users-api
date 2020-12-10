package services

import (
	"github.com/nelbermora/bookstore_users-api/domain/users"
	"github.com/nelbermora/bookstore_users-api/utils/errors"
)

var (
	UsersService userServiceInterface = &usersService{}
)

type usersService struct {
}

// esto es para el mock
type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	Update(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	FindByStatus(string) ([]users.User, *errors.RestErr)
}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
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

func (u *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	err := result.Get()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *usersService) Update(user users.User) (*users.User, *errors.RestErr) {
	/*err := user.Validate()
	if err != nil {
		return nil, err
	}*/
	result := users.User{Id: user.Id}
	err := result.Get()
	if err != nil {
		return nil, err
	}
	err = user.Update()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *usersService) DeleteUser(userId int64) *errors.RestErr {
	result := users.User{Id: userId}
	err := result.Get()
	if err != nil {
		return err
	}
	err = result.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u *usersService) FindByStatus(status string) ([]users.User, *errors.RestErr) {
	dao := users.User{}
	userList, err := dao.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return userList, nil

}
