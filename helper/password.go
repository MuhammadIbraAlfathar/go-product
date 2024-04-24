package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(generateFromPassword), err
}

func ValidatePassword(hashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	return err
}
