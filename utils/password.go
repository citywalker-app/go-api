package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"

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
	b := make([]byte, 3)
	// nolint:gosec
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	encodedStr := base64.RawStdEncoding.EncodeToString(b)
	num, err := strconv.Atoi(encodedStr[0:4])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%04d", num%9000+1000), nil
}
