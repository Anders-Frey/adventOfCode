package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	safeReports := 0

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

		if reportIsSafe(line) || reportIsKindaSafe(line) {
			fmt.Printf("Report is safe\n")
			safeReports++
		} else {
			fmt.Printf("Report is not safe!\n")
		}

		fmt.Printf("Safe reports: %d\n", safeReports)
	}
}

func reportIsSafe(report string) bool {
	reportNumbers := strings.Fields(report)
	safeRange := []int{}

	for index := range reportNumbers {
		current, _ := strconv.ParseInt(reportNumbers[index], 10, 32)
		var next int64 = 0

		if index < len(reportNumbers)-1 {
			next, _ = strconv.ParseInt(reportNumbers[index+1], 10, 32)
		} else {
			fmt.Printf("Reached end of line and still safe, phew!.\n")
			return true
		}

		// determine asc / desc in first iteration
		if index == 0 {
			if next > current {
				safeRange = []int{1, 2, 3}
			} else {
				safeRange = []int{-1, -2, -3}
			}
		}

		diff := next - current

		if diff == 0 {
			return false
		}

		if !slices.Contains(safeRange, int(diff)) {
			return false
		}
	}

	// report back with report on report safety
	// (actaully, I don't we will ever reach this point?)
	fmt.Printf("Huh, it was really possible to get to here!.\n")
	return true
}

func reportIsKindaSafe(report string) bool {
	fmt.Printf("  Report failed first test:\n")
	fmt.Printf("    Applying Problem Dampener to check if report is kinda safe\n")
	// Remove one element at the time to kinda check for safety
	reportNumbers := strings.Fields(report)
	for index, number := range reportNumbers {
		truncatedReport := append([]string(nil), reportNumbers...)
		truncatedReport = slices.Delete(truncatedReport, index, index+1)

		fmt.Printf("Checking without %s\n", number)
		if reportIsSafe(strings.Join(truncatedReport, " ")) {
			return true
		}

	}
	return false
}
