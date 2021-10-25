package aws

import (
	"context"
	"fmt"
	"io"
	"os"

	"cedugenio/go-sender-files/properties"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

func SendFileToS3(filename string, fileContent io.Reader) error {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String(properties.Props.BucketName),
		Key:    aws.String(filename),
		Body:   fileContent,
	}

	_, err = PutFile(ctx, client, input)
	if err != nil {
		fmt.Println("Got error Updating file ")
		fmt.Println(err)
	}
	return nil
}
