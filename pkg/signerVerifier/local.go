package signerVerifier

import (
	"github.com/golang-jwt/jwt/v5"
)

type Signer struct {
	key string
}

func NewSigner(key string) *Signer {
	return &Signer{
		key: key,
	}
}

func (ls *Signer) Sign(signingString string) (string, error) {
	parser := jwt.NewParser()
	signingString = signingString + "." + ls.key
	token, _, err := parser.ParseUnverified(signingString, &jwt.MapClaims{})
	if err != nil {
		return "", err
	}
	signedString, err := token.SignedString([]byte(ls.key))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func (ls *Signer) Verify(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(ls.key), nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}
	return true, nil
}
