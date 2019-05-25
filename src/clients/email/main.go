package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// The subject line for the email.
	Subject = "Amazon SES Test (AWS SDK for Go)"

	// The HTML body for the email.
	HtmlBody =  "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	//The email body for recipients with non-HTML email clients.
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// The character encoding for the email.
	CharSet = "UTF-8"
)

func handler(ctx context.Context, sesEvent events.SimpleEmailEvent) {
	for _, record := range sesEvent.Records {
		sesRecord := record.SES

		//To add in additional fields to publish to the logs add "snsRecord.'fieldname'"
		fmt.Printf("Message = %s \n", sesRecord.Mail)


		sess, err := session.NewSession(&aws.Config{
			Region:aws.String("us-west-2")},
		)
		svc := ses.New(sess)

		// Assemble the email.
		input := &ses.SendEmailInput{
			Destination: &ses.Destination{
				CcAddresses: []*string{
				},
				ToAddresses: []*string{
					aws.String(sesRecord.Mail.Source),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Html: &ses.Content{
						Charset: aws.String(CharSet),
						Data:    aws.String(HtmlBody),
					},
					Text: &ses.Content{
						Charset: aws.String(CharSet),
						Data:    aws.String(TextBody),
					},
				},
				Subject: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(Subject),
				},
			},
			Source: aws.String(sesRecord.Mail.Destination[0]),
			// Uncomment to use a configuration set
			//ConfigurationSetName: aws.String(ConfigurationSet),
		}

		// Attempt to send the email.
		result, err := svc.SendEmail(input)

		// Display error messages if they occur.
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ses.ErrCodeMessageRejected:
					fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
				case ses.ErrCodeMailFromDomainNotVerifiedException:
					fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
				case ses.ErrCodeConfigurationSetDoesNotExistException:
					fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}

			return
		}

		fmt.Println(result)
	}
}


func main() {
	lambda.Start(handler)
}