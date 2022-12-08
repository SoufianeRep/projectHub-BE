package svcs

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

var S3Client *s3.Client

func init() {
	// Load AWS bucket config and initialize S3 bucket
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// initializes a new s3 client
	S3Client = s3.NewFromConfig(cfg)
}

type PutObjectData struct {
	File   multipart.File
	Name   string
	TeamID string
}

func UploadObject(ctx *gin.Context, data PutObjectData) (*s3.PutObjectOutput, error) {
	output, err := S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String("projects/teams/" + data.TeamID + "/video/" + data.Name), // TODO
		Body:   data.File,
	})
	if err != nil {
		return &s3.PutObjectOutput{}, fmt.Errorf("could not upload the file: %v", err)
	}

	return output, nil
}
