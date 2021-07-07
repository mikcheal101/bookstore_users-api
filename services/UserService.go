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

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

/**
 * Method to get a user by Id
 */
func (UserService *UserService) GetUserById(id int64) (*users.User, *errors.RestError) {
	if id <= 0 {
		err := errors.NewBadRequestError{}
		err.BadRequest("Invalid user id")
		return nil, &err.RestError
	}

	user := users.User{Id: id}
	if userErr := user.Get(); userErr != nil {
		return nil, userErr
	}
	return &user, nil
}
