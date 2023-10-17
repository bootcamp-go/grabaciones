package tools_test

import (
	"fmt"
	"intro-unit-testing/ej2/tools"
	"testing"
)

func TestElementAtIndexWithExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := 0
	expected := 1

	// act

	result, err := tools.ElementAtIndex(slice, idx)

	// assert
	if err != nil {
		t.Fatalf("test failed. expected nil error but %v was given instead", err)
	}
	if result != expected {
		t.Fatalf("test failed. expected %d but %d was given instead", expected, result)
	}
	// if test fails, this line will never be executed
	fmt.Println("test ended")
}

func TestElementAtIndexWithNonExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := len(slice)

	// act
	_, err := tools.ElementAtIndex(slice, idx)

	// assert
	if err == nil {
		t.Fatal("test failed. expected error but nil was given instead")
	}
}
