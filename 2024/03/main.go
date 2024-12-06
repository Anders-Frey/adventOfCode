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
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// For part two we need to adhere to the do() and don't() commands.
	// Don't know if regex will still cut it, but they might look like this:
	// `do()\b` and `don't()\b`

	multiplicationPattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	multiplicationTasks := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		matches := multiplicationPattern.FindAllString(line, -1)

		for _, match := range matches {
			multiplicationTasks = append(multiplicationTasks, match)
		}

		fmt.Printf("Found %d matches on line\n", len(matches))
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
