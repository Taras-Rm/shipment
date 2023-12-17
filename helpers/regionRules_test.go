package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegionRulesFactor(t *testing.T) {
	testCases := []struct{
		name string
		code string
		expected float64
	}{
		{
			name: "nordic region country 1",
			code: "SE",
			expected: 1,
		},
		{
			name: "nordic region country 2",
			code: "DK",
			expected: 1,
		},
		{
			name: "europe country",
			code: "PL",
			expected: 1.5,
		},
		{
			name: "outside of europe country",
			code: "AE",
			expected: 2.5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := RegionRulesFactor(tC.code)

			require.Equal(t, tC.expected, actual)
		})
	}

}		