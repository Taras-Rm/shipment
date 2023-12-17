package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateEmail(t *testing.T) {
	testCases := []struct{
		name string
		email string
		err error
	}{
		{
			name: "correct email",
			email: "test@gmail.com",
			err: nil,
		},
		{
			name: "correct long email",
			email: "testtesttesttesttesttesttesttesttesttesttesttesttesttest@gmail.com",
			err: nil,
		},
		{
			name: "invalid email 1 (no @)",
			email: "testgmail.com",
			err: ErrorInvalidEmail,
		},
		{
			name: "invalid email 2 (no fist part)",
			email: "@gmail.com",
			err: ErrorInvalidEmail,
		},
		{
			name: "invalid email 2 (no last part)",
			email: "test@gmail.",
			err: ErrorInvalidEmail,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := ValidateEmail(tC.email)

			require.Equal(t, tC.err, actual)
		})
	}
}

func TestValidateName(t *testing.T) {
	testCases := []struct{
		name string
		inpName string
		err error
	}{
		{
			name: "correct name",
			inpName: "testName",
			err: nil,
		},
		{
			name: "invalid name (longer 30 chars)",
			inpName: "testNametestNametestNametestNametestNametestName",
			err: ErrorInvalidName,
		},
		{
			name: "invalid email (contain digits)",
			inpName: "test9nam8e",
			err: ErrorInvalidName,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := ValidateName(tC.inpName)

			require.Equal(t, tC.err, actual)
		})
	}
}

func TestValidateCountryCode(t *testing.T) {
	testCases := []struct{
		name string
		code string
		err error
	}{
		{
			name: "correct country code",
			code: "UA",
			err: nil,
		},
		{
			name: "invalid country code (long word)",
			code: "CODE",
			err: ErrorInvalidCountryCode,
		},
		{
			name: "invalid country code (lovercase)",
			code: "ua",
			err: ErrorInvalidCountryCode,
		},
		{
			name: "not existing country code",
			code: "UU",
			err: ErrorNotExistingCountryCode,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := ValidateCountryCode(tC.code)

			require.Equal(t, tC.err, actual)
		})
	}
}