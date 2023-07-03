package secretmanager

import (
	"encoding/json"
	"fmt"
	"github.com/acossovich/twitterGo/awsgo"
	"github.com/acossovich/twitterGo/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret

	fmt.Println(">Pido secreto" + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	_ = json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("Lectura de secret exitosa " + secretName)

	return datosSecret, nil
}
