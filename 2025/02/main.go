package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Result variable starts at 0
	result := 0

	file, err := os.Open("input_anders")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
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

			rangeStart, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("Skipping element %d: cannot parse rangeStart '%s' as integer\n", i, parts[0])
				continue
			}

			rangeEnd, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Skipping element %d: cannot parse rangeEnd '%s' as integer\n", i, parts[1])
				continue
			}

			fmt.Printf("rangeStart: %d, rangeEnd: %d\n", rangeStart, rangeEnd)

			// Loop that traverses the integers between rangeStart and rangeEnd
			for num := rangeStart; num <= rangeEnd; num++ {
				// For each integer convert it to a string
				numStr := strconv.Itoa(num)

				// Split the string in the middle
				// If the number of chars is uneven print a line and skip the entry
				if len(numStr)%2 != 0 {
					fmt.Printf("Skipping number %d: number of chars is uneven (%d chars)\n", num, len(numStr))
					continue
				}

				mid := len(numStr) / 2
				firstHalf := numStr[:mid]
				secondHalf := numStr[mid:]

				// Compare the two strings and check if they are equal
				if firstHalf == secondHalf {
					fmt.Printf("Number %d: halves are EQUAL (%s == %s)\n", num, firstHalf, secondHalf)
					// For each instance of equal strings, add the original integer to the result
					result += num
				} else {
					fmt.Printf("Number %d: halves are NOT equal (%s != %s)\n", num, firstHalf, secondHalf)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// In the end, print out the result
	fmt.Printf("\nResult: %d\n", result)
}
