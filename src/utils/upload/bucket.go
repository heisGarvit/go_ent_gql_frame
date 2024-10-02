package upload

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log/slog"
	"os"
	"project/src/infra/env"
)

var accessKey = env.Get("CLOUDFLARE_ACCESS_KEY")
var secretKey = env.Get("CLOUDFLARE_SECRET_KEY")
var endpoint = env.Get("CLOUDFLARE_ENDPOINT")
var bucket = env.Get("CLOUDFLARE_BUCKET")
var publicURL = env.Get("CLOUDFLARE_PUBLIC_URL")

func FileToBucket(
	ctx context.Context,
	uploadKey string,
	contentEncoding *string,
	filePath string,
) (string, error) {

	var err error
	// Create a new session
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(endpoint),
	})

	// Create a new uploader using the session
	uploader := s3manager.NewUploader(sess)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", filePath, err)
	}
	// Upload the file
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:          aws.String(bucket),
		Key:             aws.String(uploadKey),
		Body:            file,
		ContentEncoding: contentEncoding,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	} else {
		slog.DebugContext(ctx, "file uploaded to, ", result.Location)
	}

	return publicURL + uploadKey, nil
}
