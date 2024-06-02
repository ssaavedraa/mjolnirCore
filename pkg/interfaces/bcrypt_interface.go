package interfaces

type BcryptInterface interface {
	CompareHashAndPassword (hashedPassword, password []byte) error
}

