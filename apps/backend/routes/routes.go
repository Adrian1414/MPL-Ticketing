package routes

import (
	"net/http"

	"example.com/mpl/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	server.GET("/", getHome)

	server.GET("/teams", controllers.GetAllTeams)
	server.GET("/teams/:id", controllers.GetTeamByID)
	server.POST("/teams", controllers.CreateTeam)
	server.PUT("/teams/:id", controllers.UpdateTeam)
	server.DELETE("/teams/:id", controllers.DeleteTeam)

	server.GET("/rosters", controllers.GetAllRosters)
	server.POST("/rosters", controllers.CreateRoster)
}

func getHome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Home Page",
	})
}
