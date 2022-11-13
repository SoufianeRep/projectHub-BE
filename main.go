package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/SoufianeRep/tscit/api"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	// Load AWS bucket config and initialize S3 bucket
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Println("Load config error", err)
	}

	// initializes a new s3 client
	api.S3Client = s3.NewFromConfig(cfg)

	// initializes transcription service client
	trsSession := session.Must(session.NewSession())
	api.TrsSession = transcribeservice.New(trsSession, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}

func main() {
	server := api.NewServer()
	err := server.Start(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		log.Fatal("Cannot Start the server:", err)
	}
}
