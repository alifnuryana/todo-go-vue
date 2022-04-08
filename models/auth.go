package models

import "github.com/golang-jwt/jwt/v4"

type RequestLogin struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type CustomClaim struct {
	Username string
	Role     string
	Id       string
	jwt.RegisteredClaims
}
