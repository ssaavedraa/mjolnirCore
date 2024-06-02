package utils

import "golang.org/x/crypto/bcrypt"

type BcryptWrapper struct {}

func (b *BcryptWrapper) CompareHashAndPassword (hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}