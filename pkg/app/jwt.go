package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	c "github.com/ahmed-saleh/playbook/config"
)

type Claims struct {
	UserUuid string `json:"user_uuid"`
	jwt.StandardClaims
}

func CreateJwtAccessToken(u string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := &Claims{}
	// atClaims["authorized"] = true
	atClaims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(c.AppSetting.JwtTime)).Unix()
	atClaims.UserUuid = u

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(c.AppSetting.JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
