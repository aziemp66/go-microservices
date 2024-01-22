package password

import (
	"regexp"

	errorCommon "broker/internal/http/error"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errorCommon.NewUnauthorized("wrong credentials")
	}
	return nil
}

func PasswordValidation(password string) error {
	// Minimum 8 characters, at least one uppercase letter, one lowercase letter,
	// one digit, and one special character
	lowerCase := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !lowerCase {
		return errorCommon.NewBadRequest("password must contain at least one lowercase letter")
	}
	upperCase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !upperCase {
		return errorCommon.NewBadRequest("password must contain at least one uppercase letter")
	}
	digit := regexp.MustCompile(`\d`).MatchString(password)
	if !digit {
		return errorCommon.NewBadRequest("password must contain at least one digit")
	}
	specialChar := regexp.MustCompile(`[!@#$%^&*(){}:"|<>?\[\]\\;',./]`).MatchString(password)
	if !specialChar {
		return errorCommon.NewBadRequest("password must contain at least one special character")
	}

	length := len(password) >= 8
	if !length {
		return errorCommon.NewBadRequest("password must be at least 8 characters")
	}

	return nil
}
