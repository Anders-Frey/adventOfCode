package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "input_anders"

	res, err := Run(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\nResult: %d\n", res)
}

func Run(inputFile string) (int, error) {
	// Result variable starts at 0
	result := 0

	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("\nWorking on this line: %s\n", line)

		digit1 := 0
		digit1Pos := 0
		digit2 := 0

		// two runs through the line:
		// first run to find digit 1, and mark the cursor where digit 1 was found.
		for i := 0; i < len(line); i++ {
			valueAtCursor := int(line[i] - '0') // Convert ASCII to int8
			if valueAtCursor > digit1 && len(line) != i+1 {
				// we found a better digit1 and it was not the last digit
				digit1 = valueAtCursor
				digit1Pos = i
			}
		}

		// second run to find digit 2, starting from the cursor position of digit 1.
		for i := digit1Pos + 1; i < len(line); i++ {
			valueAtCursor := int(line[i] - '0') // Convert ASCII to int8
			if valueAtCursor > digit2 {
				// we found a better digit2
				digit2 = valueAtCursor
			}
		}

		joltage := digit1*10 + digit2
		fmt.Printf("Highest joltage on this line: %d\n", joltage)
		result += joltage
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return result, nil
}
