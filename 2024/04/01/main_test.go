package main

import (
	"testing"
)

func TestCheckForXMAS(t *testing.T) {
	test := []struct {
		chars  []rune
		result int
	}{
		{
			chars:  []rune{'X', 'M', 'A', 'S'},
			result: 1,
		},
		{
			chars:  []rune{'X', 'M', 'A', 'S', 'H'},
			result: 0,
		},
		{
			chars:  []rune{'M', 'M', 'A', 'S'},
			result: 0,
		},
	}

	// Act
	for _, test := range test {
		result := CheckForXMAS(test.chars)

		// Assert
		if test.result != result {
			t.Errorf("Expected %t but got %t", test.result, result)
		}
	}
}
