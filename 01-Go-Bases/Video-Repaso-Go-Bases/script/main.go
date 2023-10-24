package main

import "testing"

var myByte byte = 255

var booleanTrue bool = true
var booleanFalse bool = false

var firstName string = "John"
var lastName string = "Doe"

func TestHandler(t *testing.T) {
	type testCase struct {
		// name of the test case
		name string
		// set-up (arrange)
		setUpMockService func(mk *service.Mock)
		// base
		input  struct{}
		output struct{}
	}

	testCases := []testCase{}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange

			// act

			// assert
		})
	}
}