package api

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

var TrsSession *transcribeservice.TranscribeService

func TranscribeTest() (trsOutput *transcribeservice.StartTranscriptionJobOutput, err error) {
	trsOutput, err = TrsSession.StartTranscriptionJob(&transcribeservice.StartTranscriptionJobInput{
		TranscriptionJobName: aws.String("transcription_test"),
		IdentifyLanguage:     aws.Bool(true),
		MediaFormat:          aws.String("wav"),
		OutputBucketName:     aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Media: &transcribeservice.Media{
			MediaFileUri: aws.String("s3://" + os.Getenv("AWS_BUCKET_NAME") + "/gettysburg.wav"),
		},
	})

	if err != nil {
		fmt.Println(err)
		return trsOutput, err
	}

	return trsOutput, nil
}
