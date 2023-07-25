package handlers

import (
	"fmt"
"strconv"
"gambit/auth"
"github.com/aws/aws-lambda-go/events"
)


func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string){
	fmt.Println("proccess "+path+ " > " + method)
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

isOk, statusCode, user := validateAuthorization(path, method, headers)
if !isOk{
	return statusCode, user
}

switch path[0:4] {
case "user":
cas

}
	return 400, "Method Invalid"

}

func validateAuthorization( path string, method string, heeaders map[string]string)(bool, int, string) {
	if(path == "product" && method == "GET")||(path =="category" && method=="GET"){
		return true, 200, ""
	}
	token := heeaders["Authorization"]
	if len(token)==0{
		return false, 401, "Token is required"
	}
	todoOk, msg, err := auth.ValidateToken(token)
	if !todoOk{
		if err != nil{
			return false, 401, err.Error()
		}
			return false, 401, msg
	}
	fmt.Println("Token Ok")
	return true, 200,msg
}