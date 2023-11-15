package util

import (
	"gin_blog/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}

var jwtSecret []byte

func SetUp() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claim := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin_blog",
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := jwtToken.SignedString(jwtSecret)
	return token, err
}

func ParesToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
