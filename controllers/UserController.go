package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikcheal101/bookstore_users-api/domain/users"
	"github.com/mikcheal101/bookstore_users-api/services"
	"github.com/mikcheal101/bookstore_users-api/utils/errors"
)

type UserController struct{}

var userService services.UserService

func (UserController *UserController) GetUser(req *gin.Context) {
	userId, userError := strconv.ParseInt(req.Param("id"), 10, 64)
	if userError != nil {
		err := errors.NewBadRequestError{}
		err.BadRequest("Invalid User id!")
		req.JSON(err.RestError.Status, err.RestError)
		return
	}

	user, userErr := userService.GetUserById(userId)
	if userErr != nil {
		fmt.Println(userErr)
		req.JSON(userErr.Status, userErr)
		return
	}

	req.JSON(http.StatusOK, user)
}

/**
* Usercontroller method to read user
* data from a post request and create the user
 */
func (UserController *UserController) CreateUser(req *gin.Context) {
	var user users.User
	if err := req.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError{}
		message := "Invalid parameters!"
		response := restError.BadRequest(message)
		req.JSON(response.Status, response)
		return
	}
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		req.JSON(err.Status, err)
		return
	}
	req.JSON(http.StatusCreated, createdUser)
}

func (UserController *UserController) SearchUser(req *gin.Context) {
	req.String(http.StatusNotImplemented, "Pending Implementation")
}
