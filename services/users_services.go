package services

import (
	"github.com/pavangujar23/bookStore_UsersAPI/domain/users"
	"github.com/pavangujar23/bookStore_UsersAPI/utils/errorsUtil"
)

func CreateUser(user users.User) (*users.User, *errorsUtil.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errorsUtil.RestErr) {
	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
