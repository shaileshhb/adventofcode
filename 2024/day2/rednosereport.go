package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/shaileshhb/adventofcode/file"
)

func main() {
	input := file.ReadFromFile()

	reports := processFileInput(input)
	// fmt.Println(reports)

	removeArrayWithDuplicates(reports)

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
		isIncreasing := report[0] < report[1]
		isSafe := true
		for i := 1; i < len(report)-1; i++ {

			if isIncreasing {
				if report[i] < report[i+1] && (report[i+1]-report[i] >= 1) && (report[i+1]-report[i] <= 3) {
					continue
				}
				isSafe = false
				break
			}

			if report[i] > report[i+1] && (report[i]-report[i+1] >= 1) && (report[i]-report[i+1] <= 3) {
				continue
			}

			isSafe = false
			break
		}

		fmt.Println("reports", report, isSafe)
		if isSafe {
			safeReports++
		}
	}

	return safeReports
}

func removeArrayWithDuplicates(reports [][]int) {
	newReports := make([][]int, 0)

	for _, report := range reports {
		tempReport := make([]int, len(report))
		copy(tempReport, report)

		sort.Ints(tempReport)

		isDuplicateFound := false
		for i := 1; i < len(tempReport)-1; i++ {
			if tempReport[i] == tempReport[i+1] {
				isDuplicateFound = true
				break
			}
		}

		fmt.Println("== duplicate rports", tempReport, isDuplicateFound)

		if !isDuplicateFound {
			newReports = append(newReports, report)
		}
	}

	fmt.Println("len", len(newReports), len(reports))
}

func test() {
	// report := []int{7, 6, 4, 2, 1}
	report := []int{17, 18, 22}
	safeReports := 0

	isIncreasing := report[0] < report[1]
	isSafe := true
	fmt.Println("isIncreasing", isIncreasing)

	for i := 1; i < len(report)-1; i++ {
		fmt.Println("value", report[i], report[i+1])

		if isIncreasing {
			if report[i] < report[i+1] && (report[i+1]-report[i] >= 1) && (report[i+1]-report[i] <= 3) {
				continue
			}
			isSafe = false
			fmt.Println("this has to break", i, report[i])
			break
		}

		if report[i] > report[i+1] && (report[i]-report[i+1] >= 1) && (report[i]-report[i+1] <= 3) {
			continue
		}

		isSafe = false
		fmt.Println("breaking for i", i)
		break
	}

	if isSafe {
		safeReports++
	}
	fmt.Println("safe reports", safeReports)
}
