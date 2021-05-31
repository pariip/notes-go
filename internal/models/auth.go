package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pariip/notes-go/internal/models/types"
)

type Claims struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Role        types.Role `json:"role"`
	jwt.StandardClaims
}
