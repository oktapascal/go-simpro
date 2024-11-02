package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/config"
	"github.com/spf13/viper"
	"time"
)

func GenerateAccessToken(parameters *config.JwtParameters) (string, int64, error) {
	expiresAt := time.Now().Add(15 * time.Minute).Unix()

	token := config.GenerateToken(jwt.MapClaims{
		"iss":         viper.GetString("APP_NAME"),
		"sub":         parameters.Username,
		"aud":         parameters.Role,
		"exp":         expiresAt,
		"iat":         time.Now().Unix(),
		"id":          parameters.Id,
		"menu_group":  parameters.GroupMenu,
		"permissions": parameters.Permissions,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SIGNATURE_KEY")))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt, nil
}

func GenerateRefreshToken(parameters *config.JwtParameters) (string, int64, error) {
	expiresAt := time.Now().Add(7 * (24 * time.Hour)).Unix()

	token := config.GenerateToken(jwt.MapClaims{
		"iss":         viper.GetString("APP_NAME"),
		"sub":         parameters.Username,
		"aud":         parameters.Role,
		"exp":         expiresAt,
		"iat":         time.Now().Unix(),
		"id":          parameters.Id,
		"menu_group":  parameters.GroupMenu,
		"permissions": parameters.Permissions,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_REFRESH_SIGNATURE_KEY")))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt, nil
}

func VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := config.VerifyToken(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(viper.GetString("JWT_SIGNATURE_KEY")), nil
	})

	return token, err
}

func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := config.VerifyToken(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(viper.GetString("JWT_REFRESH_SIGNATURE_KEY")), nil
	})

	return token, err
}
