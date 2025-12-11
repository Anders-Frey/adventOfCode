package main

import (
	"os"
	"testing"
)

func TestExampleSum(t *testing.T) {
	// Write the example input into a temporary file named like `input_anders`
	data := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	tmp := "tmp_file"
	if err := os.WriteFile(tmp, []byte(data), 0o644); err != nil {
		t.Fatalf("failed to write input file: %v", err)
	}
	defer os.Remove(tmp)

	got, err := Run(tmp)
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	// Expected sum for invalid IDs in the example
	var want int64 = 4174379265
	if got != want {
		t.Fatalf("unexpected sum: got %d want %d", got, want)
	}
}
