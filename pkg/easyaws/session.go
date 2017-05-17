package easyaws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func SessionFromEnvVars() *session.Session {
	return session.Must(
		session.NewSessionWithOptions(
			session.Options{
				Config: aws.Config{
					Credentials: credentials.NewEnvCredentials(),
				},
			},
		),
	)
}

func SessionFromProfile(Profile string) *session.Session {
	return session.Must(
		session.NewSessionWithOptions(
			session.Options{
				Profile: Profile,
				SharedConfigState: session.SharedConfigEnable,
			},
		),
	)
}

func SessionFromStaticCreds(AwsAccessKeyID string, AwsSecretAccessKey string, Region string) *session.Session {
	return session.Must(
		session.NewSessionWithOptions(
			session.Options{
				Config: aws.Config{
					Credentials: credentials.NewStaticCredentials(AwsAccessKeyID, AwsSecretAccessKey, ""),
					Region: &Region,
				},
			},
		),
	)
}
