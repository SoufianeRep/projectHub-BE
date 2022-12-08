package main

import (
	"log"
	"os"

	"github.com/SoufianeRep/tscit/api"
	"github.com/SoufianeRep/tscit/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
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

	// initializes transcription service client
	trsSession := session.Must(session.NewSession())
	api.TrsSession = transcribeservice.New(trsSession, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}
