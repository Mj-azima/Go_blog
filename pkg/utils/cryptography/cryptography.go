package cryptography

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	passwd, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return nil, err
	}
	return passwd, nil
}

func CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}
