package main

import (
	"testing"
)

func TestExampleSum(t *testing.T) {
	got, err := Run("input_test")
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	// Expected sum for invalid IDs in the example
	var want int = 3121910778619
	if got != want {
		t.Fatalf("unexpected sum: got %d want %d", got, want)
	}
}
