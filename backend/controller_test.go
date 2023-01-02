package backend

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	num1 := 5
	num2 := 10

	result1 := AddNumbers(num1, num2)
	result2 := 5 + 10
	require.Equal(t, result1, result2)
}
