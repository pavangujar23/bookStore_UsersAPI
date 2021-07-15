package users

import (
	"fmt"

	"github.com/pavangujar23/bookStore_UsersAPI/utils/errorsUtil"
)

var (
	usersDb = make(map[int64]*User)
)

//Work with pointer to modify
func (user *User) Get() *errorsUtil.RestErr {
	result := usersDb[user.Id]

	if result == nil {
		return errorsUtil.NewUserNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errorsUtil.RestErr {
	current := usersDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errorsUtil.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errorsUtil.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}
	usersDb[user.Id] = user

	return nil
}
