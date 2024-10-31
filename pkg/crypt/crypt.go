package crypt

import "golang.org/x/crypto/bcrypt"

type Crypt struct{}

func NewCrypt() *Crypt {
	return &Crypt{}
}

func (c *Crypt) GenerateHash(target string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(target), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (c *Crypt) CompareHash(hashed string, target string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(target))
	return err == nil
}
