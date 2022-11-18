package main

import (
	"context"
	"log"
	"os"

	"github.com/SoufianeRep/tscit/api"
	"github.com/SoufianeRep/tscit/db"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Load AWS bucket config and initialize S3 bucket
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Get a db instance
	db := db.GetDB()

	server, err := api.NewServer(db)
	if err != nil {
		log.Fatal("cannot create a new server:", err)
	}

	err = server.Start(os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		log.Fatal("Cannot start the server:", err)
	}

	// initializes a new s3 client
	api.S3Client = s3.NewFromConfig(cfg)

	// initializes transcription service client
	trsSession := session.Must(session.NewSession())
	api.TrsSession = transcribeservice.New(trsSession, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}
