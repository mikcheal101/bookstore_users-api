package app

import (
	"github.com/mikcheal101/bookstore_users-api/controllers"
)

type Routes struct{}

var pingController controllers.PingController = controllers.PingController{}
var userController controllers.UserController = controllers.UserController{}

func (Routes *Routes) MapUrls() {
	router.GET("/ping", pingController.PingRequest)

	router.GET("/users/:id/", userController.GetUser)
	router.GET("/users/search/", userController.SearchUser)
	router.POST("/users/", userController.CreateUser)
}
