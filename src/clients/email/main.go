package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"os"
	"strings"

	"github.com/jhillyerd/enmime"
    "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	Subject = "Diplomacy Engine"
	HtmlBody =  "<h1>Command received</p>"
	TextBody = "Command received"

	NEHtmlBody = "<h1>Requested game name was not found</h1></br>Please use another, or contact your game administrator to create new one"
	NETextBody = "Requested game name was not found\nPlease use another, or contact your game administrator to create new one"
)

type Move struct {
	Player string
	Turn int
	Command string
}

func sendMail(sess *session.Session, subject string, html string, text string, from string, to string) {
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
			},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(html),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(text),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(from),
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

func handler(ctx context.Context, sesEvent events.SimpleEmailEvent) {
	for _, record := range sesEvent.Records {
		sesRecord := record.SES

		sess, err := session.NewSession(&aws.Config{
			Region:aws.String("eu-west-1")},
		)

		fmt.Printf("MessageID = %s \n", sesRecord.Mail.MessageID)

		s3svc := s3.New(sess)

		out, err := s3svc.GetObject(&s3.GetObjectInput {
			Bucket: aws.String("diplomacy-mails"),
			Key: aws.String(sesRecord.Mail.MessageID),
		})
		if err != nil {
			fmt.Println("Cannot access s3 object")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Parse message body with enmime.
		env, err := enmime.ReadEnvelope(out.Body)
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Printf("Message text = %s \n", env.Text)
		fmt.Printf("Message html = %s \n", env.HTML)

		dynamoSvc := dynamodb.New(sess)

		item := &Move{
			Turn: 1,
			Player: sesRecord.Mail.Source,
			Command: env.Text,
		}

		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			fmt.Println("Request not parsed properly")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		gameDomain := strings.Split(sesRecord.Mail.Destination[0],"@")

		tables, err := dynamoSvc.ListTables(&dynamodb.ListTablesInput{})
		if err != nil {
			fmt.Println("Cannot list DynamoDB tables:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		gameExists := false
		for _, table := range tables.TableNames {
			if gameDomain[0] == *table {
				gameExists = true
				break
			}
		}

		if !gameExists {
			sendMail(sess, Subject, NEHtmlBody, NETextBody, "no-reply@diplomacy.mbien.pl", sesRecord.Mail.Source)
			continue
		}

		dynamoInput := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(gameDomain[0]),
		}

		_, err = dynamoSvc.PutItem(dynamoInput)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}


		sendMail(sess, Subject + " - game " + gameDomain[0], HtmlBody, TextBody, sesRecord.Mail.Destination[0], sesRecord.Mail.Source)
	}
}


func main() {
	lambda.Start(handler)
}
