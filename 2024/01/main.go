package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sortedLeft := []int{}
	sortedRight := []int{}
	sumDistance := 0

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("\nWorking on this line %s\n", line)
		temp := strings.Fields(line)
		tempLeft, _ := strconv.ParseInt(temp[0], 10, 32)
		tempRight, _ := strconv.ParseInt(temp[1], 10, 32)

		sortedLeft = append(sortedLeft, int(tempLeft))
		sortedRight = append(sortedRight, int(tempRight))
	}

	sort.Ints(sortedLeft)
	sort.Ints(sortedRight)

	for i := 0; i < len(sortedLeft); i++ {
		diff := sortedRight[i] - sortedLeft[i]
		fmt.Printf("         Diff: %d\n", diff)
		float64Diff := float64(diff)
		unsignedDiff := math.Abs(float64Diff)
		fmt.Printf("Unsigned Diff: %f\n", unsignedDiff)
		backToIntDiff := int(unsignedDiff)
		fmt.Printf("Back to int  : %d\n", backToIntDiff)

		sumDistance += backToIntDiff
	}

	// Answer to part 1
	fmt.Printf("Sum of distances: %d\n", sumDistance)

	// Start on part 2
	similarityScore := 0

	for i := 0; i < len(sortedLeft); i++ {
		leftDigit := sortedLeft[i]
		occurenceCounter := 0
		fmt.Printf("Left: %d\n", sortedLeft[i])
		for i := 0; i < len(sortedLeft); i++ {
			if leftDigit == sortedRight[i] {
				occurenceCounter++
			}
		}
		fmt.Printf("Occurence: %d\n", occurenceCounter)

		similarityScore += (occurenceCounter * leftDigit)
		fmt.Printf("Similarity Score: %d\n", similarityScore)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
