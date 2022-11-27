package interfaces

type ICryptography interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword []byte, password []byte) error
}
