package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/luchonicolosi/gambituser/awsgo"
	"github.com/luchonicolosi/gambituser/db"
	"github.com/luchonicolosi/gambituser/models"
)

func main() {
	lambda.Start(EjecutoLamda)
}

func EjecutoLamda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAws()
	
	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)

		case "sub":
			datos.UserUUID = att
			fmt.Println("Id =" + datos.UserUUID)
		}
	}

	err := db.SignUp(datos)

	return event, err
}

