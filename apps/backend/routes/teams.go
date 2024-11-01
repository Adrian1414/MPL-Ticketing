package routes

import (
	"net/http"

	"example.com/mpl/models"
	"github.com/gin-gonic/gin"
)

func getAllTeams(context *gin.Context) {
	teams, err := models.GetAllTeams()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get all teams",
		})
		return
	}
	context.JSON(http.StatusOK, teams)
}
