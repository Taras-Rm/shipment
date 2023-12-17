package helpers

import (
	"errors"
	"regexp"

	"github.com/biter777/countries"
)

 var (
	ErrorInvalidEmail error = errors.New("invalid email")
	ErrorInvalidName error = errors.New("invalid email")
	ErrorInvalidAddress error = errors.New("invalid address")
	ErrorInvalidCountryCode error = errors.New("invalid country code")
	ErrorNotExistingCountryCode error = errors.New("not existing country code")
 )

// validate email
func ValidateEmail(email string) error {
	var emailReg = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailReg.MatchString(email) {
		return ErrorInvalidEmail
	}

	return nil
}

// validate name
func ValidateName(name string) error {
	var nameReg = regexp.MustCompile("^[^0-9]+$")

	if len(name) > 30 || !nameReg.MatchString(name) {
		return ErrorInvalidName
	}

	return nil
}

// validate country code
func ValidateCountryCode(code string) error {
	var codeReg = regexp.MustCompile("^[A-Z]{2}$")

	if !codeReg.MatchString(code) {
		return ErrorInvalidCountryCode
	}

	// check if there is a country with such a code
	country := countries.ByName(code).String()
	if country == "Unknown" {
		return ErrorNotExistingCountryCode
	}

	return nil
}

// validate address
func ValidateAddress(address string) error {
	var addressReg = regexp.MustCompile("^[a-zA-Z\\s]+\\s\\d+\\,\\s[a-zA-Z0-9\\s]+\\s[0-9]+$")

	if len(address) > 100 || !addressReg.MatchString(address) {
		return ErrorInvalidAddress
	}

	return nil
}
