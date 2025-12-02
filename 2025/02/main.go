package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
