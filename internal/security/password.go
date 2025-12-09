package security

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyPassword = errors.New("password is empty")
	ErrLengthPassword = errors.New("password most be 8 character")
)


func HashPassword(password string) (string, error) {

	if len(password) == 0 {
		return "", ErrEmptyPassword
	}

	if len(password) < 8 {
		return "", ErrLengthPassword
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt generate failed : %w", err)
	}
	
	return string(bytes), err
}
