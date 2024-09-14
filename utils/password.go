package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func IsBcryptHash(s string) bool {
	if len(s) != 60 {
		return false
	}

	matched, err := regexp.MatchString(`^\$2[aby]\$.{56}$`, s)
	if err != nil {
		return false
	}

	return matched
}

func GeneratePassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashedPwd, plainPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)) == nil
}

func GenerateRandomCode() (string, error) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// nolint:gosec
	// Generate a random number between 0 and 9999
	code := rand.Intn(10000)

	// Format the number as a 4-digit string, padding with zeros if necessary
	return fmt.Sprintf("%04d", code), nil
}
