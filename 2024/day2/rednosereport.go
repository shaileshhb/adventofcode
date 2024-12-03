package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shaileshhb/adventofcode/file"
)

func main() {
	input := file.ReadFromFile()

	reports := processFileInput(input)
	// fmt.Println(reports)

	safeReports := getSafeReportCount(reports)
	fmt.Println("safe reports: ", safeReports)
	// test()
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
		if isReportSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func isReportSafe(report []int) bool {
	if report[0] == report[1] {
		return false
	}

	isIncreasing := report[0] < report[1]
	diff := int(math.Abs(float64(report[0] - report[1])))

	if diff < 0 || diff > 3 {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		diff := int(math.Abs(float64(report[i] - report[i+1])))

		if isIncreasing {
			if report[i] < report[i+1] && (diff <= 3) {
				continue
			}
			return false
		}

		if report[i] > report[i+1] && (diff <= 3) {
			continue
		}
		return false
	}

	return true
}

// func reportDampener() {
// 	// tolerate a single bad level

// 	report := []int{1, 3, 2, 4, 5}

// }
