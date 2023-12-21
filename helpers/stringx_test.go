package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StrToPointerStr(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		inputStr := "test string"

		actual := StrToPointerStr(inputStr)

		require.Equal(t, &inputStr, actual)
	})
	
}