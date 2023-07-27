package handlers

import (
	"fmt"
	"gambit/auth"
	"gambit/routers"
	"strconv"

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
	return UserProcess(body, path, method, user, id, request)
case "prod":
	return ProductProcess(body, path, method, user, idn, request)
case "stoc":
	return StockProcess(body, path, method, user, idn, request)
case "addr":
	return AddressProcess(body, path, method, user, idn, request)
case "cate":
	return CategoryProcess(body, path, method, user, idn, request)
case "orde":
	return OrderProcess(body, path, method, user, idn, request)
}
	return 400, "Method Invalid"

}
func UserProcess(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest)(int, string){
	return 400, "Method Invalid"
}
func ProductProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest)(int, string){
	return 400, "Method Invalid"
}
func CategoryProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest)(int, string){
	switch method {
		case "GET":
		return routers.SelectCategories(body, request)
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(body, user, id)
	}

	return 400, "Method Invalid"
}

func StockProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest)(int, string){
	return 400, "Method Invalid"
}
func AddressProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest)(int, string){
	return 400, "Method Invalid"
}

func OrderProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest)(int, string){
	return 400, "Method Invalid"
}

func validateAuthorization( path string, method string, headers map[string]string)(bool, int, string) {
	if(path == "product" && method == "GET")||(path =="category" && method=="GET"){
		return true, 200, ""
	}
	token := headers["authorization"]
	if len(token)==0{
		return false, 401, "Token is required"
	}
	fmt.Println(headers)
	fmt.Println("token  > "+token)
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