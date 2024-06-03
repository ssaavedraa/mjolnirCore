package interfaces

type BcryptInterface interface {
	CompareHashAndPassword (hashedPassword, password []byte) error
	GenerateFromPassword (password []byte, cost int) ([]byte, error)
}

