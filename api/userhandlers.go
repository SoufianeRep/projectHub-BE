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
}

func handleCreateUser(ctx *gin.Context) {
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
		// TODO: Handle properly dev pusposes only!!!
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "an Error has occured while hashing the passwor",
		})
		return
	}

	arg := db.CreateUserParams{
		Email:      req.Email,
		Password:   hashedPassword,
		LastSignin: time.Now(),
	}

	err = db.CreateUser(arg)
	if err != nil {
		fmt.Println("creating error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

type createUserWithTeamRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	TeamName string `json:"teamName" binding:"required"`
}

type userLoginParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func handleLogin(ctx *gin.Context) {
	var req userLoginParams

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
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid credentials."})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wring.",
			})
			return
		}
	}

	if err := util.CheckPassword(user.Password, req.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User authenticated",
	})
}
