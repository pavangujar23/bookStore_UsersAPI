package users

import (
	"strings"

	"github.com/pavangujar23/bookStore_UsersAPI/utils/errorsUtil"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errorsUtil.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errorsUtil.NewBadRequestError("invalid email address")
	}
	return nil
}
