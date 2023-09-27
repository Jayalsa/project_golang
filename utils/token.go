package utils

import (
	"fmt"
	"jayalsa/project_golang/constants"
	"jayalsa/project_golang/entities"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["email"] = email
	tokenString, err := token.SignedString(constants.SecretKey)
	return tokenString, err
}

func DecodeToken(tokenInput string) bool {
	token, err := jwt.Parse(tokenInput, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
	})

	if err == nil && token.Valid {
		fmt.Println("Your token is valid.  I like your style.")

	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
		return false
	}
	return token.Valid
}

func GenerateAllTokens(email string, firstName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &entities.SignedDetails{
		Email:      email,
		First_name: firstName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &entities.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(constants.SecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(constants.SecretKey))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
