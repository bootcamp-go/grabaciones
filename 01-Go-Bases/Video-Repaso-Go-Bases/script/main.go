package main

import "testing"

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
}