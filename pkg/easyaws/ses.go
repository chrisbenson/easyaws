package easyaws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendMail(to []string, from string, cc []string, bcc []string, subject string, body string, awsSession *session.Session) (string, error) {
	svc := ses.New(awsSession)
	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			BccAddresses: aws.StringSlice(bcc),
			CcAddresses:  aws.StringSlice(cc),
			ToAddresses:  aws.StringSlice(to),
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data: aws.String(body), // Required
					//Charset: easyaws.String("Charset"),
				},
				Text: &ses.Content{
					Data: aws.String("MessageData"), // Required
					//Charset: easyaws.String("Charset"),
				},
			},
			Subject: &ses.Content{ // Required
				Data: aws.String(subject), // Required
				//Charset: easyaws.String("Charset"),
			},
		},
		Source: aws.String(from), // Required
		//ConfigurationSetName: easyaws.String("ConfigurationSetName"),
		ReplyToAddresses: []*string{
			aws.String(from), // Required
			// More values...
		},
		ReturnPath: aws.String(from),
		//ReturnPathArn: easyaws.String("AmazonResourceName"),
		//SourceArn:     easyaws.String("AmazonResourceName"),
		//Tags: []*ses.MessageTag{
		//	{ // Required
		//		Name:  easyaws.String("MessageTagName"),  // Required
		//		Value: easyaws.String("MessageTagValue"), // Required
		//	},
		// More values...
		//},
	}
	output, err := svc.SendEmail(params)
	if err != nil {
		return "", err
	}
	return *output.MessageId, nil
}
