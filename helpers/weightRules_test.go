package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWeightClassRulesFactor(t *testing.T) {
	testCases := []struct{
		name string
		weight float64
		result uint
	}{
		{
			name: "small weight class",
			weight: 2,
			result: 100,
		},
		{
			name: "small weight class (left border)",
			weight: 1,
			result: 100,
		},
		{
			name: "small weight class (right border)",
			weight: 10,
			result: 100,
		},
		{
			name: "medium weight class",
			weight: 20,
			result: 300,
		},
		{
			name: "medium weight class (left border)",
			weight: 11,
			result: 300,
		},
		{
			name: "medium weight class (right border)",
			weight: 25,
			result: 300,
		},
		{
			name: "large weight class",
			weight: 35,
			result: 500,
		},
		{
			name: "large weight class (left border)",
			weight: 26,
			result: 500,
		},
		{
			name: "large weight class (right border)",
			weight: 50,
			result: 500,
		},
		{
			name: "huge weight class",
			weight: 700,
			result: 2000,
		},
		{
			name: "huge weight class (left border)",
			weight: 51,
			result: 2000,
		},
		{
			name: "huge weight class (right border)",
			weight: 1000,
			result: 2000,
		},
		{
			name: "not supported weight class 1",
			weight: 0,
			result: 0,
		},
		{
			name: "not supported weight class 2",
			weight: 1100,
			result: 0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := WeightClassRulesFactor(tC.weight)

			require.Equal(t, tC.result, actual)
		})
	}
}