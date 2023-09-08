package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var secretAccesKey = []byte("aswyTgHUI122JgJNDdjsi900Hg)*&")
var secretRefreshKey = []byte("kjdb&bsj)khd89Bhsnau5GHJnd")

func GenerateAccesTokens(claims *jwt.MapClaims) (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accesToken, err := tokens.SignedString(secretAccesKey)
	if err != nil {
		return "", err
	}

	return accesToken, nil
}

func GenerateRefreshTokens(claims *jwt.MapClaims) (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := tokens.SignedString(secretRefreshKey)

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func VerifyAccesToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return secretAccesKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return secretRefreshKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeRefreshTokens(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyRefreshToken(tokenString)
	if err != nil {
		return nil, err
	}

	//isOk merupakan variabel yang menampung data apakah valid atau tidak
	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
