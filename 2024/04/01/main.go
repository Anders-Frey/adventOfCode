package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	// Open file and create the 2d slice
	// Turns out it is quadratic. 140 X 140
	file, err := os.Open("./../input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var inputAsSlice [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		inputAsSlice = append(inputAsSlice, row)
	}

	// Traverse all directions and sum up # XMAS
	sum := 0
	sum += Horizontal(inputAsSlice)
	sum += Vertical(inputAsSlice)
	sum += DiagonalRightUp(inputAsSlice)
	sum += DiagonalRightDown(inputAsSlice)

	fmt.Printf("Found %d XMAS\n", sum)
}

func Horizontal(input [][]rune) int {
	fmt.Printf("Looking Horizontal - both ways\n")
	sum := 0

	for indexX := 0; indexX < len(input); indexX++ {
		for indexY := 0; indexY < len(input); indexY++ {
			fmt.Printf("Doing %d,%d\n", indexX, indexY)
			sequenceToCheck := input[indexX][indexY : indexY+4]

			fmt.Printf("Checking sequence: %c\n", sequenceToCheck)
			sum += CheckForXMAS(sequenceToCheck)
		}
	}
	fmt.Printf("Horizontal found #%d XMAS\n", sum)
	return sum
}

func Vertical(input [][]rune) int {
	fmt.Printf("Looking Vertical - both ways\n")
	sum := 0

	for indexX := 0; indexX < len(input); indexX++ {
		for indexY := 0; indexY < len(input); indexY++ {
			fmt.Printf("Doing %d,%d\n", indexX, indexY)
			if indexY+3 >= len(input) {
				fmt.Printf("Reached end at %d,%d\n", indexX, indexY)
				continue
			}
			sequenceToCheck := []rune{
				input[indexY][indexX],
				input[indexY+1][indexX],
				input[indexY+2][indexX],
				input[indexY+3][indexX],
			}
			fmt.Printf("Checking sequence: %c\n", sequenceToCheck)
			sum += CheckForXMAS(sequenceToCheck)
		}
	}
	fmt.Printf("Vertical found #%d XMAS\n", sum)
	return sum
}

func DiagonalRightUp(input [][]rune) int {
	fmt.Printf("Looking Diagonal Right Up - both ways\n")
	sum := 0

	for indexX := 3; indexX < len(input); indexX++ { // Don't start at top, there's no room
		for indexY := 0; indexY < len(input); indexY++ {
			fmt.Printf("Doing %d,%d\n", indexX, indexY)
			if indexY+3 >= len(input) {
				fmt.Printf("Reached end at %d,%d\n", indexX, indexY)
				continue
			}
			sequenceToCheck := []rune{
				input[indexX][indexY],
				input[indexX-1][indexY+1],
				input[indexX-2][indexY+2],
				input[indexX-3][indexY+3],
			}
			fmt.Printf("Checking sequence: %c\n", sequenceToCheck)
			sum += CheckForXMAS(sequenceToCheck)
		}
	}
	fmt.Printf("Diagonal Right Up found #%d XMAS\n", sum)
	return sum
}

func DiagonalRightDown(input [][]rune) int {
	fmt.Printf("Looking Diagonal Right Down - both ways\n")
	sum := 0

	for indexX := 0; indexX < len(input)-3; indexX++ {
		for indexY := 0; indexY < len(input); indexY++ {
			fmt.Printf("Doing %d,%d\n", indexX, indexY)
			if indexY+3 >= len(input) {
				fmt.Printf("Reached end at %d,%d\n", indexX, indexY)
				continue
			}
			sequenceToCheck := []rune{
				input[indexX][indexY],
				input[indexX+1][indexY+1],
				input[indexX+2][indexY+2],
				input[indexX+3][indexY+3],
			}
			fmt.Printf("Checking sequence: %c\n", sequenceToCheck)
			sum += CheckForXMAS(sequenceToCheck)
		}
	}
	fmt.Printf("Diagonal Right Up found #%d XMAS\n", sum)
	return sum
}

func CheckForXMAS(chars []rune) int {
	matchesFound := 0

	if len(chars) != 4 {
		return matchesFound
	}

	// Remember to check both ways
	if slices.Compare(chars, []rune{'X', 'M', 'A', 'S'}) == 0 {
		fmt.Printf("  Found XMAS!!\n")
		matchesFound++
	}
	if slices.Compare(chars, []rune{'S', 'A', 'M', 'X'}) == 0 {
		fmt.Printf("  Found SAMX!!\n")
		matchesFound++
	}

	return matchesFound
}
