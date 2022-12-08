package svcs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

var TrsSession *transcribeservice.TranscribeService

func init() {
	session := session.Must(session.NewSession())
	TrsSession = transcribeservice.New(session, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}
