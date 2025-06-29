package helpers

import (
	"crypto/rand"
	"kutamukti-api/pkg/exceptions"
	"math/big"
	"net/http"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

func GeneratePassword(length int) (string, *exceptions.Exception) {
	if length <= 0 {
		return "", nil
	}

	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", exceptions.NewException(http.StatusInternalServerError, exceptions.ErrCredentialsHash)
		}
		password[i] = charset[num.Int64()]
	}

	return string(password), nil
}
