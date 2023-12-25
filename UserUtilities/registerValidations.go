package UserUtilities

import (
	"net/mail"
	"strings"
)

func IsValidUsername(username string) bool {
	if len(username) < MIN_USERNAME_LENGTH || len(username) > MAX_USERNAME_LENGTH {
		return false
	}
	for _, char := range username {
		if !strings.ContainsRune(VALID_USERNAME_CHARS, char) {
			return false
		}
	}
	return true
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil || len(email) == 0
}

func IsValidSecrets(secretQuestion string, secretAnswer string) bool {
	if len(secretQuestion) < MIN_SECRET_QUESTION_LENGTH || len(secretQuestion) > MAX_SECRET_QUESTION_LENGTH {
		return false
	}
	if !(len(secretAnswer) == 0 || len(secretAnswer) == 64) {
		return false
	}
	if len(secretAnswer) == 64 && len(secretQuestion) == 0 {
		return false
	}
	if len(secretQuestion) > 0 && len(secretAnswer) == 0 {
		return false
	}
	return true
}

func IsValidPassword(password string) bool {
	if len(password) != 64 {
		return false
	}
	for _, char := range password {
		if char < '0' || char > '9' {
			if char < 'a' || char > 'f' {
				return false
			}
		}
	}
	return true
}
