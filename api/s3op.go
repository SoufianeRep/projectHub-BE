package api

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

var S3Client *s3.Client

// UploadObject Uploads file to aws S3 bucket an return an error
func UploadObject(ctx *gin.Context, object *multipart.FileHeader) (err error) {
	file, err := object.Open()
	if err != nil {
		panic("Could not open the file")
	}

	_, err = S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(object.Filename),
		Body:   file,
	})

	TranscribeTest()

	file.Close()
	if err != nil {
		fmt.Println("Could not upload the file", err)
		return err
	}

	return nil
}
