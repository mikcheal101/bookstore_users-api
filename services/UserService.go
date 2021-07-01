package services

import "github.com/mikcheal101/bookstore_users-api/domain/users"

type UserService struct{}

/**
* Method to create a user
 */
func (UserService *UserService) CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
