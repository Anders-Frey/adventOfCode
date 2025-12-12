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

		// TODO
		// Now we need to find the 12 highest digits on the line. There's always 100 digits pr line.
		// At the beginning all 100 digits are eligible candiates
		// We count the number of 9's first, then 8's etc.
		// Then we sum from 9 and downwards until we have 12.
		// All candidates with digits below this are ignored.

	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return result, nil
}
