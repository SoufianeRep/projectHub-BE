package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SoufianeRep/tscit/db"
	"github.com/SoufianeRep/tscit/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	TeamName string `json:"team_name" binding:"omitempty,min=2"`
}

type userResponse struct {
	ID         uint   `json:"id"`
	Email      string `json:"email"`
	LastSignin time.Time
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:         user.ID,
		Email:      user.Email,
		LastSignin: user.LastSignin,
	}
}

func (server *Server) handleCreateUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid informations",
		})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		// TODO: Handle properly dev purposes only!!!
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "an error has occured while hashing the password",
		})
		return
	}

	arg := db.CreateUserParams{
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := db.CreateUser(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	team := db.Team{}
	if req.TeamName != "" {
		tArg := db.CreateTeamParams{
			TeamName: req.TeamName,
		}

		team, err := db.CreateTeam(tArg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		rArg := db.CreateRoleParams{
			UserID: user.ID,
			TeamID: team.ID,
			Role:   "manager",
		}

		err = db.CreateRole(rArg)
		if err != nil {
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
	}

	res := newUserResponse(user)
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
		"team": team,
	})
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	User        userResponse `json:"user"`
	AccessToken string       `json:"access_token"`
}

func (server *Server) handleLogin(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials.",
		})
		return
	}

	user, err := db.GetUserByEmail(req.Email)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "no such user exists"})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong.",
			})
			return
		}
	}

	if err := util.CheckPassword(user.Password, req.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials.",
		})
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.ID,
		user.Email,
		time.Minute*15, // TODO: change the validity of the token to a env variable for global use
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	res := loginUserResponse{
		User:        newUserResponse(user),
		AccessToken: accessToken,
	}

	// res := loginUserResponse{}
	ctx.JSON(http.StatusOK, res)
}

type getUserRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

func handleGetUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	user, err := db.GetUSerByID(req.ID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "no such user exists"})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong.",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
