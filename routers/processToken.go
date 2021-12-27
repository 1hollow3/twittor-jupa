package routers

import (
	"errors"
	"strings"

	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email is the value of email to use in all the endpoints
var Email string

// IDUser is the id returned to the model, which one will be used in all the endpoints
var IDUser string

// ProcessToken is to extract the token values
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("MAX JuanPeace LAN")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format is invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := db.UserExist(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

}
