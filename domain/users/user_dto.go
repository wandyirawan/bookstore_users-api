package users

import (
	"strings"

	"github.com/wandyirawan/bookstore_users-api/utils/errors"
)

type User struct {
	Id         int64  `json:"id,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	DateCreted string `json:"date_creted,omitempty"`
}

func (user *User) Validete() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
