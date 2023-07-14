package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gihub.com/04-gambit/user-gambit/awsgo"
	"gihub.com/04-gambit/user-gambit/bd"
	"gihub.com/04-gambit/user-gambit/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(StartLambda)
}

func StartLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()
	if !ValidoParametros() {
		fmt.Println("error en los parámetros, debe enviar Secret Manager")
		err := errors.New("error en elos parámetros debe enviar 'SecretManager'")
		return event, err

	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)

		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}

	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret" + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)

	return event, err

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")

	return traeParametro
}
