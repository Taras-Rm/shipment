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