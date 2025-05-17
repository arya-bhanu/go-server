package utils

import "golang.org/x/crypto/bcrypt"

func HashPasssword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), 10)
}

func IsPasswordValid(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
