package users

import (
	"strings"

	"github.com/mikcheal101/bookstore_users-api/utils/errors"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Method to validate user email
// @return error RestError | nil
func (user *User) ValidateUser() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		error := errors.NewBadRequestError{}
		return error.BadRequest("invalid  email address")
	}
	return nil
}
