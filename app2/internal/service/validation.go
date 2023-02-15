package service

import "regexp"

func ValidateEmail(email string) bool {
	emailRegex := `^[a-z0-9]+@[a-z.-]+\.[a-z]{2,}$`
	pattern := regexp.MustCompile(emailRegex)
	return pattern.MatchString(email)
}
