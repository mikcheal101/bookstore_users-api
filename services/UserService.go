package services

import (
	"github.com/mikcheal101/bookstore_users-api/domain/users"
	"github.com/mikcheal101/bookstore_users-api/utils/errors"
)

type UserService struct{}

/**
* Method to create a user
 */
func (UserService *UserService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	return nil, nil
}
