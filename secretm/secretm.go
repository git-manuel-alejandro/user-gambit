package secretm

import (
	"encoding/json"
	"fmt"

	"gihub.com/04-gambit/user-gambit/awsgo"
	"gihub.com/04-gambit/user-gambit/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJSON, error) {
	var datosSecret models.SecretRDSJSON
	fmt.Println(" > Pido secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())

		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secret OK" + nombreSecret)

	return datosSecret, nil

}
