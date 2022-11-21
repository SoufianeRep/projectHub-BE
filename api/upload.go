package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("unable to reade file headers", err)
		return
	}

	ct := strings.Split(file.Header.Values("Content-Type")[0], "/")
	if ct[0] != "video" && ct[0] != "audio" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong file type",
		})
	}

	err = UploadObject(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not upload the fileplease try again.",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Upload successful.",
	})
}
