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
	var nameReg = regexp.MustCompile("^[^0-9]+$")

	if len(name) > 30 || !nameReg.MatchString(name) {
		return errors.New("invalid name")
	}

	return nil
}

// validate country code
func ValidateCountryCode(code string) error {
	var codeReg = regexp.MustCompile("^[A-Z]{2}$")

	if !codeReg.MatchString(code) {
		return errors.New("invalid country code")
	}

	return nil
}

// validate country code
func ValidateAddress(address string) error {
	var addressReg = regexp.MustCompile("^[a-zA-Z\\s]+\\s\\d+\\,\\s[a-zA-Z0-9]+\\s[0-9]+$")

	if len(address) > 100 || !addressReg.MatchString(address) {
		return errors.New("invalid address")
	}

	return nil
}
