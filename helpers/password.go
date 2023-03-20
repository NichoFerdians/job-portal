package helpers

import (
	"errors"
	"fmt"

	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckPasswordMD5(password string, hashedPassword string) error {
	hash := md5.Sum([]byte(password))

	if hex.EncodeToString(hash[:]) != hashedPassword {
		return errors.New("Invalid credentials")
	}

	return nil
}
