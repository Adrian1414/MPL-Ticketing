package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	server.GET("/", getHome)
	server.GET("/teams", getAllTeams)
}

func getHome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Home Page",
	})
}
