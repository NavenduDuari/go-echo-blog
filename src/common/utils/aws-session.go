package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetSnsSession() *session.Session {
	return session.New(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("GO_ECHO_BLOG_AWS_SNS_ACCESS_KEY"), os.Getenv("GO_ECHO_BLOG_AWS_SNS_SECRET_KEY"), ""),
	})
}
