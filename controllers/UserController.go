package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController *UserController) GetUser(req *gin.Context) {
	req.String(http.StatusNotImplemented, "Pending Implementation")
}

func (UserController *UserController) CreateUser(req *gin.Context) {
	req.String(http.StatusNotImplemented, "Pending Implementation")
}

func (UserController *UserController) SearchUser(req *gin.Context) {
	req.String(http.StatusNotImplemented, "Pending Implementation")
}
