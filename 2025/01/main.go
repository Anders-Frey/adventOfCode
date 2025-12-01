package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dialPosition := 50
	zeroCounter := 0
	clickCounter := 0

	file, err := os.Open("input_anders")
	// file, err := os.Open("den_gir_16.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		clicks := 0
		startDialPosition := dialPosition
		line := scanner.Text()
		fmt.Printf("\nWorking on this line: %s\n", line)
		fmt.Printf("Current dial position: %d\n", startDialPosition)
		direction := string(line[0])
		amount, _ := strconv.Atoi(line[1:])

		// Start by extracting number of full rounds
		numberOfRounds := (amount / 100) % 10
		fmt.Printf("Amount: %d, rounds: %d\n", amount, numberOfRounds)

		// Then subtract the rounds from the amount and continue with this easier to handle number
		amount = amount - (numberOfRounds * 100)
		fmt.Printf("Reduced amount after rounds: %d\n", amount)

		// Now look for clicks when passing 0 or 100
		// Do not count clicks if we start at 0. The zeroCounter handles those
		if direction == "R" {
			dialPosition += amount
			if dialPosition > 100 && startDialPosition != 0 {
				fmt.Printf("Dial position exceeded 100: %d\n", dialPosition)
				clicks++ // click when passing 100
			}
		} else {
			dialPosition -= amount
			if dialPosition < 0 && startDialPosition != 0 {
				fmt.Printf("Dial position below 0: %d\n", dialPosition)
				clicks++ // click when passing 0
			}
		}

		fmt.Printf("numberOfRounds and Clicks this round: %d, %d\n", numberOfRounds, clicks)
		clickCounter += numberOfRounds + clicks
		fmt.Printf("Total clicks so far: %d\n", clickCounter)

		if dialPosition >= 0 {
			dialPosition = dialPosition % 100
		} else {
			dialPosition = 100 + dialPosition
		}

		fmt.Printf("New dial position: %d\n", dialPosition)

		if dialPosition == 0 {
			zeroCounter++
			fmt.Printf("Dial is 0. Zero counter at: %d\n", zeroCounter)
		}

	}

	fmt.Printf("Number of times dial hit zero: %d\n", zeroCounter)
	fmt.Printf("Number of clicks: %d\n", clickCounter)
	fmt.Printf("Total: %d\n", zeroCounter+clickCounter)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
