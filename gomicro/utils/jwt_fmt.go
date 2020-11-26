package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("jey.sign")


type MyCustomClaims struct {
	UserName string `json:"user_name"`
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func GetToken(username string,userid int64)(tokenStr string,err error){
	claims := MyCustomClaims{
		username,
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
			Issuer:    "bg",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenStr, err = token.SignedString(mySigningKey); err != nil {
		BgLogger.Errorf("get token failed err :",err)
		return "",err
	}
	return
}

func ValidateToken(tokenStr string)(userid int64,username string,err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("token过期")
				return
			} else {
				fmt.Println(err)
				return
			}
		}
		return
	}
	//获取用户信息(100w次 损耗0.6s)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = claims["user_name"].(string)
		userid = int64(claims["user_id"].(float64))
	} else {
		return
	}
	return
}