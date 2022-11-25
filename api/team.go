package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SoufianeRep/tscit/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

type getTeamRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type getTeamResponse struct {
	ID       uint         `json:"id"`
	TeamName string       `json:"team_name"`
	Projects []db.Project `json:"projects"`
	Members  []db.User    `json:"team_members"`
}

func handleGetTeam(ctx *gin.Context) {
	var req getTeamRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid team ID",
		})
		return
	}

	team, err := db.GetTeam(req.ID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "no such team exist",
			})
			return
		default:
			fmt.Println("Error:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wront",
			})
		}
	}

	projects, err := team.GetProjects()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	members, err := team.GetMembers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	res := getTeamResponse{
		ID:       team.ID,
		TeamName: team.TeamName,
		Projects: projects,
		Members:  members,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
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
		return
	}

	req.TeamID = uint(teamID)

	// TODO: continue witht he logic
}
