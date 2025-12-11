package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input_anders"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	res, err := Run(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\nResult: %d\n", res)
}

// isRepeatedSequence returns true when s is composed of some non-empty
// digit sequence repeated at least twice to form the whole string.
// Run processes the provided input file and returns the sum of invalid IDs.
func Run(inputFile string) (int64, error) {
	// Result variable starts at 0
	var result int64 = 0

	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("\nWorking on this line: %s\n", line)

		// Split the line by comma separator into a slice
		contents := strings.Split(line, ",")
		fmt.Printf("Number of elements: %d\n", len(contents))
		fmt.Printf("Contents: %v\n", contents)

		// Loop that enumerates each element in contents
		for i, element := range contents {
			fmt.Printf("\nProcessing element %d: %s\n", i, element)

			// For each element extract two integers separated by a dash into variables called rangeStart and rangeEnd
			parts := strings.Split(element, "-")

			// If the format of the element does not allow this, print a line with info about that and skip the element
			if len(parts) != 2 {
				fmt.Printf("Skipping element %d: format does not contain exactly one dash separator\n", i)
				continue
			}

			rangeStart, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				fmt.Printf("Skipping element %d: cannot parse rangeStart '%s' as integer\n", i, parts[0])
				continue
			}

			rangeEnd, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				fmt.Printf("Skipping element %d: cannot parse rangeEnd '%s' as integer\n", i, parts[1])
				continue
			}

			fmt.Printf("rangeStart: %d, rangeEnd: %d\n", rangeStart, rangeEnd)

			// Loop that traverses the integers between rangeStart and rangeEnd
			for num := rangeStart; num <= rangeEnd; num++ {
				// For each integer convert it to a string
				numStr := strconv.FormatInt(num, 10)

				// Check if the ID is invalid by being a repeated digit-sequence
				if isRepeatedSequence(numStr) {
					fmt.Printf("Number %d: is INVALID (repeated sequence)\n", num)
					result += num
				} else {
					fmt.Printf("Number %d: is valid\n", num)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return result, nil
}

func isRepeatedSequence(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}

	// Try every possible block length from 1 up to n/2
	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}
		repeats := n / l
		if repeats < 2 {
			continue
		}

		block := s[:l]
		ok := true
		for i := 1; i < repeats; i++ {
			if s[i*l:(i+1)*l] != block {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	return false
}
