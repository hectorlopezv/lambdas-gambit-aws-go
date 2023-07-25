package auth

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenJSON struct {
	Sub       string `json:"sub"`
	Event_Id  string `json:"event_id"`
	Token_use string 	`json:"token_use"`
	Scope     string 	`json:"scope"`
	Auth_time int 		`json:"auth_time"`
	Iss       string 	`json:"iss"`
	Exp       int 		`json:"exp"`
	Iat       int 		`json:"iat"`
	Client_id string    `json:"client_id"`
	Username  string    		`json:"cognito:username"`
}

func ValidateToken(token string) (bool, string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, "el token no es valido", fmt.Errorf("invalid token")
	}
    parsedToken, err := jwt.Parse(token, nil)
    if parsedToken == nil {
        return false, err.Error(), err
    }
    claims, ok := parsedToken.Claims.(jwt.MapClaims)

	fmt.Println("claims > ")
	fmt.Println(claims)
	fmt.Println("parts > ")
	fmt.Println(parts)
	if !ok  {
		fmt.Println("Error en traer el token", err.Error())
		return false, err.Error(), err
	}
	jsonData, err := json.Marshal(claims)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return false, err.Error(), err
	}

	// Print the JSON data (optional)
	fmt.Println("JSON Data:", string(jsonData))

	// Create a new struct instance
	var tokenData TokenJSON

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, &tokenData)
	if err != nil {
		fmt.Println("Error en Unmarshal", err.Error())
		return false, err.Error(), err
	}
	fmt.Println("tokenData > ")
	fmt.Println(tokenData)
	now := time.Now()
	tm := time.Unix(int64(tokenData.Exp), 0)
	if tm.Before(now) {
		return false, "el token expiro", fmt.Errorf("token expired")
	}
	return true, string(tokenData.Username), nil
}
