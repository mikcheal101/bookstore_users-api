package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikcheal101/bookstore_users-api/domain/users"
	"github.com/mikcheal101/bookstore_users-api/services"
)

type UserController struct{}

var userService services.UserService

func (UserController *UserController) GetUser(req *gin.Context) {
	req.String(http.StatusNotImplemented, "Pending Implementation")
}

/**
* Usercontroller method to read user
* data from a post request and create the user
 */
func (UserController *UserController) CreateUser(req *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(req.Request.Body)
	// if err != nil {
	// 	//TODO: Handle error
	// 	fmt.Println("Invalid Parameters")
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO: Handle json error
	// 	fmt.Println("Invalid Parameters")
	// 	return
	// }
	if err := req.ShouldBindJSON(&user); err != nil {
		//TODO: return bad request error
		req.JSON(http.StatusBadRequest, "Invalid parameters!")
		return
	}
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		//TODO: Handle user service creation error
		fmt.Println(err)
		return
	}
	fmt.Println(user, createdUser)
	req.JSON(http.StatusCreated, createdUser)
}

func (UserController *UserController) SearchUser(req *gin.Context) {
	req.String(http.StatusNotImplemented, "Pending Implementation")
}
