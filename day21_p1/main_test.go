package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		startA int
		startB int
		answer int
	}{
		{4, 8, 739785},
	}

	for _, test := range tests {
		result := solve(test.startA, test.startB)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
