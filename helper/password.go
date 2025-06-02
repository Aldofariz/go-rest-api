package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

func VerifyPassword(hashpassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err
}

