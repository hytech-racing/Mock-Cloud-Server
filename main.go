package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hytech-racing/Mock-Cloud-Server/app"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file %s", err)
	// }

	// Setup aws s3 connection
	aws_region := os.Getenv("AWS_REGION")
	if aws_region == "" {
		log.Fatal("could not get aws region environment variable")
	}

	aws_bucket := os.Getenv("AWS_S3_RUN_BUCKET")
	if aws_region == "" {
		log.Fatal("could not get aws run bucket environment variable")
	}

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	if awsAccessKey == "" {
		log.Fatal("could not get aws access key environment variable")
	}

	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	if awsSecretKey == "" {
		log.Fatal("could not get aws secret key environment variable")
	}

	// We are creating one connection to AWS S3 and passing that around to all the methods to save resources
	s3_respository := app.NewS3Session(awsAccessKey, awsSecretKey, aws_region, aws_bucket)

	// s3Repo := &app.S3Repository{}
	a := app.New(s3_respository)

	err := a.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}

// add docker file, documentation, clean up mock data
