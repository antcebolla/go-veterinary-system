package utils

import "regexp"

func IsValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := regexp.MustCompile(`^(1\s?)?((\d{3})|\(\d{3}\))\s?-?\s?\d{3}-?\d{4}$`)
	return phoneRegex.MatchString(phoneNumber)
}
