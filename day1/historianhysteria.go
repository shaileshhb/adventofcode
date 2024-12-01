package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputLeft := make([]int, 0, 0)
	inputRight := make([]int, 0, 0)
	readFromFile(&inputLeft, &inputRight)

	distanceArr := make([]int, 0, len(inputLeft))

	sort.Ints(inputLeft)
	sort.Ints(inputRight)

	// making an assumption here that both arrays are of same length
	for i := range inputLeft {
		distance := inputLeft[i] - inputRight[i]
		distance = int(math.Abs(float64(distance)))
		distanceArr = append(distanceArr, distance)
	}

	finalDistance := 0

	for _, distance := range distanceArr {
		finalDistance += distance
	}

	fmt.Println("Distance: ", finalDistance)
}

func readFromFile(leftArr *[]int, rightArr *[]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Print each line
		line := strings.TrimSpace(scanner.Text())
		// Compile a regular expression to match one or more spaces
		re := regexp.MustCompile(`\s+`)

		// Replace multiple spaces with a single space
		line = re.ReplaceAllString(line, " ")

		lineArr := strings.Split(line, " ")
		num, err := strconv.Atoi(lineArr[0])
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return
		}

		*leftArr = append(*leftArr, num)

		num, err = strconv.Atoi(lineArr[1])
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return
		}
		*rightArr = append(*rightArr, num)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
	}
}
