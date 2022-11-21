package services

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
}

func (b *Bcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	passwd, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return nil, err
	}
	return passwd, nil
}

func (b *Bcrypt) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}
