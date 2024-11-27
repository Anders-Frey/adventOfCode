package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	sum := 0
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("\nWorking on this line %s\n", line)
		digitSlice := []string{}
		for _, r := range line {
			if unicode.IsDigit(r) {
				// replace kind of appends and therefore grows the slice,
				// but at least we overwrite the second element so that
				// when the iteration is finished, we have our first and last
				// digit in the first to indicies of the slice.
				if len(digitSlice) > 0 {
					digitSlice = slices.Replace(digitSlice, 1, 1, string(r))
				} else {
					digitSlice = append(digitSlice, string(r))
				}
			}
		}

		// for single digit lines, replicate the single digit
		if len(digitSlice) == 1 {
			digitSlice = append(digitSlice, digitSlice[0])
		}

		// cut off tail, if any
		fmt.Printf("Sorted digits %s\n", digitSlice)
		if len(digitSlice) > 1 {
			digitSlice = digitSlice[0:2]
		}

		// make to int and due the sum up
		digit, _ := strconv.Atoi(strings.Join(digitSlice, ""))
		fmt.Printf("Adding %d to %d\n", digit, sum)
		sum += digit
		fmt.Printf("  New sum is: %d\n", sum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
