package ex_9

import "regexp"

func CheckValidPassword(password string) bool {
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	if !lowercaseRegex.MatchString(password) {
		return false
	}

	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(password) {
		return false
	}

	digitRegex := regexp.MustCompile(`[0-9]`)
	if !digitRegex.MatchString(password) {
		return false
	}

	specialCharRegex := regexp.MustCompile(`[$#@]`)
	if !specialCharRegex.MatchString(password) {
		return false
	}

	if len(password) < 6 || len(password) > 12 {
		return false
	}

	return true
}
