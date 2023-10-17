package tools_test

import (
	"intro-unit-testing/ej2/tools"
	"testing"
)

func TestElementAtIndex_ExistentIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 0

	expected := 1

	// act
	result, err := tools.ElementAtIndex(slice, index)

	// assert
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
	t.Log("success")
}

func TestElementAtIndex_NonExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := 4

	expected := 0

	// act
	result, err := tools.ElementAtIndex(slice, idx)

	// assert
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if err != tools.ErrIndexOutOfRange {
		t.Fatalf("Expected %s, got %s", tools.ErrIndexOutOfRange.Error(), err.Error())
	}
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
