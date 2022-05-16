package auth

import (
	"golang.org/x/crypto/bcrypt"
)

/* -------------------- compare pwd from db and login pwd ------------------- */
func ComparePassword(input, dbPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(input))
	return err
}

/* -------------------------------- hash pwd -------------------------------- */
func GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
