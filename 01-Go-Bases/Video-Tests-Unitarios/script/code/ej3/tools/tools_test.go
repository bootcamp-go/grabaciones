package tools_test

import (
	"intro-unit-testing/ej3/tools"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElementAtIndexWithExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := 0
	expected := 1

	// act
	result, err := tools.ElementAtIndex(slice, idx)

	// assert
	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestElementAtIndexWithNonExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := len(slice)

	// act
	_, err := tools.ElementAtIndex(slice, idx)

	// assert
	require.ErrorIs(t, err, tools.ErrIndexOutOfBoundaries)
}
