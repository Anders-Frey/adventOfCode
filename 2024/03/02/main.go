package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./../input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// For part two we need to adhere to the do() and don't() commands.
	// Pattern for part #2
	multiplicationPattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)

	multiplicationTasks := []string{}
	do := true // at the begininng mul is enabled
	isMultiplicationTask := true

	for scanner.Scan() {
		line := scanner.Text()
		matches := multiplicationPattern.FindAllString(line, -1)
		doCounter := 0
		dontCounter := 0

		for _, match := range matches {
			isMultiplicationTask = true
			if match == "do()" {
				do = true
				isMultiplicationTask = false
				doCounter++
			}
			if match == "don't()" {
				do = false
				isMultiplicationTask = false
				dontCounter++
			}

			if do && isMultiplicationTask {
				multiplicationTasks = append(multiplicationTasks, match)
			}
		}

		fmt.Printf("Found %d matches on line\n", len(matches))
		fmt.Printf("Found %d   do's on line\n", doCounter)
		fmt.Printf("Found %d dont's on line\n", dontCounter)
	}

	fmt.Printf("Found %d multiplication tasks\n", len(multiplicationTasks))

	sum := 0

	for _, task := range multiplicationTasks {
		sum += multiply(task)
	}
	fmt.Printf("Sum is %d\n", sum)
}

func multiply(input string) int {
	tempString := strings.ReplaceAll(input, "mul(", "")
	tempString = strings.ReplaceAll(tempString, ")", "")
	numbers := strings.Split(tempString, ",")
	number1, _ := strconv.Atoi(numbers[0])
	number2, _ := strconv.Atoi(numbers[1])
	return number1 * number2
}
