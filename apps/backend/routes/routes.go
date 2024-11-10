package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	server.GET("/", getHome)

	server.GET("/teams", getAllTeams)
	server.GET("/teams/:id", getTeamByID)
	server.POST("/teams", createTeam)
	server.PUT("/teams/:id", updateTeam)
	server.DELETE("/teams/:id", deleteTeam)

	server.GET("/rosters", getAllRosters)
	server.POST("/rosters", createRoster)
}

func getHome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Home Page",
	})
}
