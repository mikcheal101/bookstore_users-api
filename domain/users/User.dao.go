package users

import (
	"fmt"

	"github.com/mikcheal101/bookstore_users-api/utils/errors"
)

// temp or sample data
var usersDB = make(map[int64]*User)

// Method to save the user to the database
func (user *User) Save() *errors.RestError {
	var currentUser *User = usersDB[user.Id]
	if currentUser != nil {
		var error = errors.NewBadRequestError{}
		if currentUser.Email == user.Email {
			error.BadRequest(fmt.Sprintf("[+] User with email address: %s, already exists!", user.Email))
		}
		error.BadRequest(fmt.Sprintf("[+] User with id: %d, already exists!", user.Id))
		return &error.RestError
	}

	// create the user
	usersDB[user.Id] = user
	return nil
}

// Method to get the user from the database by Id
func (user *User) Get() *errors.RestError {
	var result = usersDB[user.Id]
	if result == nil {
		var responseError = errors.NewBadRequestError{}
		return responseError.BadRequest(fmt.Sprintf("User with id: %d not found", user.Id))
	}

	// set the user data
	user.Id = result.Id
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.UpdatedAt = result.UpdatedAt
	return nil
}
