package jwt

import (
	"time"

	"github.com/1hollow3/twittor-jupa/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MAX JuanPeace LAN")

	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.Name,
		"lastname": t.Lastname,
		"birthday": t.BirthDay,
		"bio":      t.Bio,
		"location": t.Location,
		"website":  t.WebSite,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
