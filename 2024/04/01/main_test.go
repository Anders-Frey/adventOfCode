package main

import (
	"testing"
)

/*
The test_input looks like this and contains 18 "XMAS"
|MMMSXXMASM| |....XXMAS.|
|MSAMXMSMSA| |.SAMXMS...|
|AMXSXMAAMM| |...S..A...|
|MSAMASMSMX| |..A.A.MS.X|
|XMASAMXAMM| |XMASAMX.MM|
|XXAMMXXAMA| |X.....XA.A|
|SMSMSASXSS| |S.S.S.S.SS|
|SAXAMASAAA| |.A.A.A.A.A|
|MAMMMXMMMM| |..M.M.M.MM|
|MXMXAXMASX| |.X.X.XMASX|

Horizontal right     3
Horizontal left      2
Vertical down        1
Vertical up          2
Diagonal right up    4
Diagonal right down  1
Diagonal left up     4
Diagonal left down   1
SUM                 18
*/

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
