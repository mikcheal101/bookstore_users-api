package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func StartApplication() {
	var routes Routes = Routes{}
	routes.MapUrls()
	router.Run(":8080")
}
