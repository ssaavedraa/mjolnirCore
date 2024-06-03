package utils

import "golang.org/x/crypto/bcrypt"

type BcryptWrapper struct {}

func (b *BcryptWrapper) CompareHashAndPassword (hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (b *BcryptWrapper) GenerateFromPassword (password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}