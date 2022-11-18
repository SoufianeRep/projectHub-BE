package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SoufianeRep/tscit/db"
	"github.com/gin-gonic/gin"
)

type createTeamRequest struct {
	TeamName string `json:"teamName" binding:"required"`
}

func handleCreateTeam(ctx *gin.Context) {
	var req createTeamRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid information",
		})
		return
	}

	arg := db.CreateTeamParams{
		TeamName: req.TeamName,
	}

	team, err := db.CreateTeam(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "An Error has occured while creating the team",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Team Successfuly Created",
		"data":    gin.H{"teamName": team.TeamName},
	})
}

type addMemberToTeamRequest struct {
	TeamID uint
	Email  string `json:"email" binding:"required,email"`
	Role   string `json:"role" binding:"required"`
}

func handleAddMemberToTeam(ctx *gin.Context) {
	var req addMemberToTeamRequest
	teamID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("cant convert param to int")
	}
	req.TeamID = uint(teamID)

	// TODO: continue witht he logic
}
