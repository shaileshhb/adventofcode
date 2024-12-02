package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shaileshhb/adventofcode/file"
)

func main() {
	input := file.ReadFromFile()

	reports := processFileInput(input)
	fmt.Println(reports)

	safeReports := getSafeReportCount(reports)
	fmt.Println("safe reports: ", safeReports)
}

func processFileInput(input [][]string) [][]int {
	reports := make([][]int, 0, len(input))

	for _, data := range input {
		numArr := strings.Split(data[0], " ")
		arr := make([]int, 0, len(numArr))

		for _, n := range numArr {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}

			arr = append(arr, num)
		}
		reports = append(reports, arr)
	}

	return reports
}

func getSafeReportCount(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		isIncreasing := false
		previousNum := report[0]
		for i, num := range report {
			if i == 0 {
				continue
			}

			if isIncreasing && previousNum <= num {
				isIncreasing = false
				break
			}

		}
	}

	return safeReports
}
