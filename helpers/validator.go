package helpers

import (
	"errors"
	"regexp"
)

// validate email
func ValidateEmail(email string) error {
	var emailReg = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailReg.MatchString(email) {
		return errors.New("invalid email")
	}

	return nil
}

// validate name
func ValidateName(name string) error {
	var nameRegex = regexp.MustCompile("^[^0-9]+$")

	if len(name) > 30 || !nameRegex.MatchString(name) {
		return errors.New("invalid name")
	}

	return nil
}

// validate country code
func ValidateCountryCode(code string) error {
	var nameRegex = regexp.MustCompile("^[A-Z]{2}$")

	if !nameRegex.MatchString(code) {
		return errors.New("invalid country code")
	}

	return nil
}
