package users

import (
	"fmt"
	"net/mail"
	"strings"

	"example.com/hello/github.com/miftahulrespati/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Validate() *errors.RestErr { // Validate method not function
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if len(user.Password) < 8 {
		fmt.Println(len(user.Password))
		return errors.NewBadRequestError("invalid password, must at least 8 characters")
	}

	return nil
}
