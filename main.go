package main

import (
	"context"
	"gambit/awsgo"
	"gambit/db"
	"gambit/handlers"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)
func main() {
	lambda.Start(LambdaHandler)
}
func LambdaHandler(ctx context.Context, request events.APIGatewayV2HTTPRequest )(*events.APIGatewayProxyResponse, error) {
	awsgo.InicializaAws()
	if !ValidateParams(){
		panic("Error al validar parametros")
	}
	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	headers := request.Headers
	db.ReadSecret()	
	statusCode, msg := handlers.Handlers(path, method, body, headers, request)

	headerRes := map[string]string{
		"Content-Type": "application/json",
	}
	res = &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headerRes,
		Body:       string(msg),
	}

	return res, nil
}

func ValidateParams() bool{
	_, ok := os.LookupEnv("SecretName")
	if !ok {
		return ok
	}
	_, ok = os.LookupEnv("Region")
	if !ok {
		return ok
	}
	_, ok = os.LookupEnv("UserPoolId")
	if !ok {
		return ok
	}
	_, ok = os.LookupEnv("UrlPrefix")
	if !ok {
		return ok
	}
	return ok

}