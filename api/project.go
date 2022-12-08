package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/SoufianeRep/tscit/db"
	"github.com/SoufianeRep/tscit/svcs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type projectResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Language string `json:"langauge"`
	Length   uint   `json:"length"`
	// Transcript pgtype.JSONB `json:"transcript"`
}

func ProjectResponse(project db.Project) projectResponse {
	return projectResponse{
		ID:       project.ID,
		Name:     project.Name,
		Language: project.Language,
		Length:   project.Length,
		// Transcript: project.Transcript,
	}
}

type projectRequestForm struct {
	Name   string                `form:"name"`
	File   *multipart.FileHeader `form:"file" binding:"required"`
	TeamID uint
}

func handleCreateProject(ctx *gin.Context) {
	var req projectRequestForm

	object, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("unable to reade file headers", err)
		return
	}

	ct := strings.Split(object.Header.Values("Content-Type")[0], "/")
	if ct[0] != "video" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong file type",
		})
		return
	}

	file, err := object.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	defer file.Close()

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid form data",
		})
		return
	}

	od := svcs.PutObjectData{
		File:   file,
		Name:   req.Name,
		TeamID: ctx.Param("id"),
	}

	output, err := svcs.UploadObject(ctx, od)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	fmt.Println(output)

	// teamID, err := strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	fmt.Println("cant convert param to int")
	// 	return
	// }
	// req.TeamID = uint(teamID)

	// arg := db.CreateProjectParams{
	// 	Name:   req.Name,
	// 	TeamID: req.TeamID,
	// }

	// project, err := db.CreateProject(arg)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "An error has occured while creating the project",
	// 	})
	// 	return
	// }

	// res := projectResponse{
	// 	ID:   project.ID,
	// 	Name: project.Name,
	// 	// Length:     project.Length,
	// 	// Language:   project.Language,
	// 	// Transcript: project.Transcript,
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Project created successfully",
	// 	"data":    res,
	// })
}

func handleGetProject(ctx *gin.Context) {
	pi, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	projectID := uint(pi)
	project, err := db.GetProject(projectID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "no such project exist",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
		}
	}

	res := projectResponse{
		ID:       project.ID,
		Name:     project.Name,
		Length:   project.Length,
		Language: project.Language,
		// Transcript: project.Transcript,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
