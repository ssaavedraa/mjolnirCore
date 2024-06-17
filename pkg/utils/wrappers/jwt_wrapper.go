package utils

import (
	"hex/mjolnir-core/pkg/interfaces"

	"github.com/golang-jwt/jwt/v5"
)

type JwtWrapper struct{}

func (j *JwtWrapper) NewWithClaims(
	method jwt.SigningMethod,
	claime jwt.Claims,
) interfaces.JwtTokenInterface {
	return jwt.NewWithClaims(method, claime)
}

type JwtTokenWrapper struct {
	*jwt.Token
}

func (t *JwtTokenWrapper) SignedString(key interface{}) (string, error) {
	return t.Token.SignedString(key)
}
