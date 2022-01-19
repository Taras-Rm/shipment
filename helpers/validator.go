package helpers

import (
	"errors"
	"net/mail"
	"regexp"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	return nil
}

func ValidateName(name string) error {
	var nameRegex = regexp.MustCompile("[^0-9]")

	if len(name) > 30 || !nameRegex.MatchString(name) {
		return errors.New("very long name")
	}

	return nil
}
