package interfaces

import "github.com/golang-jwt/jwt/v5"

type JwtInterface interface {
	NewWithClaims (method jwt.SigningMethod, claims jwt.Claims) JwtTokenInterface
}

type JwtTokenInterface interface {
	SignedString (key interface{}) (string, error)
}