package secretm

import (
	"encoding/json"
	"gambit/awsgo"
	"gambit/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)



func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
	}
	result, err := svc.GetSecretValue(awsgo.Ctx, input)
	if err != nil {
		println("Error en GetSecretValue", err.Error())
		return secretData, err
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}
	


	err = json.Unmarshal([]byte(secretString), &secretData)
	if err != nil {
		println("Error en Unmarshal", err.Error())
		return secretData, err
	}

	return secretData,nil
}
