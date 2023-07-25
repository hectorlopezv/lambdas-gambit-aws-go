package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidateToken(token string) (bool, string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, "el token no es valido", fmt.Errorf("invalid token")
	}
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Error en base64", err.Error())
		return false, err.Error(), err
	}
	var tokenJson TokenJSON
	err = json.Unmarshal(userInfo, &tokenJson)
	if err != nil {
		fmt.Println("Error en Unmarshal", err.Error())
		return false, err.Error(), err
	}
	now := time.Now()
	tm := time.Unix(int64(tokenJson.Exp), 0)
	if tm.Before(now) {
		return false, "el token expiro", fmt.Errorf("token expired")
	}
	return true, string(tokenJson.Username), nil
}
