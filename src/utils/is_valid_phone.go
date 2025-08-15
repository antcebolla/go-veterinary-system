package utils

import "regexp"

func IsValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := regexp.MustCompile(`^\+?[\d\s\-\(\)]{7,20}$`)
	return phoneRegex.MatchString(phoneNumber)
}
