package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (PingController *PingController) PingRequest(req *gin.Context) {
	req.String(http.StatusOK, "PONG")
}
