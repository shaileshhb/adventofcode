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
	readAndProcessFromFile(&inputLeft, &inputRight)

	sort.Ints(inputLeft)
	sort.Ints(inputRight)

	distance := getDistance(inputLeft, inputRight)

	fmt.Println("Distance: ", distance)

	similarityScore := getSimilarity(inputLeft, inputRight)
	fmt.Println("similarityScore: ", similarityScore)
}

func readAndProcessFromFile(inputLeft, inputRight *[]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
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
			panic(err)
		}

		*inputLeft = append(*inputLeft, num)

		num, err = strconv.Atoi(lineArr[1])
		if err != nil {
			panic(err)
		}
		*inputRight = append(*inputRight, num)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getDistance(inputLeft, inputRight []int) int {
	distanceArr := make([]int, 0, len(inputLeft))

	for i := range inputLeft {
		distance := inputLeft[i] - inputRight[i]
		distance = int(math.Abs(float64(distance)))
		distanceArr = append(distanceArr, distance)
	}

	distance := 0

	for _, d := range distanceArr {
		distance += d
	}

	return distance
}

func getSimilarity(inputLeft, inputRight []int) (similarityScore int) {
	// left list: 3 and right list it is repeated 3 times -> 3*3 = 9
	similarityMap := make(map[int]int)

	for i := range inputLeft {
		for j := range inputRight {
			if inputLeft[i] == inputRight[j] {
				similarityMap[inputLeft[i]] += 1
				continue
			}

			if inputRight[j] > inputLeft[i] {
				break
			}
		}
	}

	for key, value := range similarityMap {
		similarityScore += key * value
	}

	return
}
