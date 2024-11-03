package routes

import (
	"net/http"
	"strconv"
	// "strings"

	"example.com/mpl/models"
	"github.com/gin-gonic/gin"
)

func getAllTeams(c *gin.Context) {
	teams, err := models.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get all teams",
		})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func getTeamByID(c *gin.Context) {
	teamId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid team ID",
		})
		return
	}
	team, err := models.GetTeamByID(teamId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get team",
		})
		return
	}
	c.JSON(http.StatusOK, team)
	}

func createTeam(c *gin.Context) {
	var team models.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	err = team.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create team",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Team created successfully",
		"team":    team,
	})
}

func updateTeam(c *gin.Context) {
	teamId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// teamName := c.Param("name")
	// teamName = strings.ToUpper(teamName)
	if err != nil {
	// if teamName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid team ID",
		})
		return
	}
	var team models.Team
	err = c.ShouldBindJSON(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	team.ID = teamId
	// team.Name = teamName
	err = team.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update team",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Team updated successfully",
		"team":    team,
	})
}

func deleteTeam(c *gin.Context) {
	teamId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid team ID",
		})
		return
	}
	team, err := models.GetTeamByID(teamId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get team",
		})
		return
	}
	err = team.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete team",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Team deleted successfully",
	})
}
