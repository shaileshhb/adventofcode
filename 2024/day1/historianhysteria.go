package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/shaileshhb/adventofcode/file"
)

func main() {
	data := file.ReadFromFile()

	inputLeft := make([]int, 0)
	inputRight := make([]int, 0)
	processFileInput(&inputLeft, &inputRight, data)

	sort.Ints(inputLeft)
	sort.Ints(inputRight)

	distance := getTotalDistance(inputLeft, inputRight)
	fmt.Println("Distance: ", distance)

	similarityScore := getSimilarityScore(inputLeft, inputRight)
	fmt.Println("similarityScore: ", similarityScore)
}

func processFileInput(inputLeft, inputRight *[]int, data [][]string) {
	for _, v := range data {
		lineArr := strings.Split(v[0], " ")

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
}

func getTotalDistance(inputLeft, inputRight []int) int {
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

func getSimilarityScore(inputLeft, inputRight []int) int {
	similarityScore := 0
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

	return similarityScore
}
