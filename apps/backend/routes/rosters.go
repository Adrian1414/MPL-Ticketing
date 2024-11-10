package routes

import (
	"net/http"

	"example.com/mpl/models"
	"github.com/gin-gonic/gin"
)

func getAllRosters(c *gin.Context) {
	rosters, err := models.GetAllRosters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get all teams",
		})
		return
	}
	c.JSON(http.StatusOK, rosters)
}

func createRoster(c *gin.Context) {
	var roster models.Roster
	err := c.ShouldBindJSON(&roster)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	err = roster.SaveRoster()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create roster",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Team created successfully",
		"roster":    roster,
	})
}