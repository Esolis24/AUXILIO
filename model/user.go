package model

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nombre         string    `json:"Nombre"`
	Password       string    `json:"Password"`
	Email          string    `json:"Email"`
	Identificacion string    `json:"Identificacion"`
	Pais           string    `json:"Pais"`
	FecNacimiento  time.Time `json:"FecNacimiento"`
}

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

// GetAuthToken returns the auth token to be used
func (u *User) GetAuthToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(tokenSecret)
	return authToken, err
}
