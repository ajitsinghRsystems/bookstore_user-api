package users

import (
	"bookstore_user-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"DateCreated"`
}

func (user *User) Validate() *errors.RestErr{
user.Email= strings.TrimSpace(strings.ToLower(user.Email))
if user.Email==""{
	return errors.NewBadRequestError("Invalid email address")

}
return nil
}