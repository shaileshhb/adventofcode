package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/shaileshhb/adventofcode/file"
)

// 30916125 -> too low
func main() {
	input := file.ReadFromFile()

	text := ""

	for _, t := range input {
		text += t[0]
	}

	mulMatches := getMullStrings(text)
	score := getTotalScore(mulMatches)
	fmt.Println("score: ", score)
}

func getMullStrings(text string) []string {
	// Compile the regex
	re := regexp.MustCompile(`mul\(\d+,\s*\d+\)`)

	// Find all matches
	matches := re.FindAllString(text, -1)

	return matches
}

func getTotalScore(mulMatches []string) int {
	numRegex := regexp.MustCompile(`\d+`)
	score := 0

	for _, match := range mulMatches {
		numbers := numRegex.FindAllString(match, -1)
		fmt.Printf("For '%s': %v\n", match, numbers)
		score += (convertStringToNum(numbers[0]) * convertStringToNum(numbers[1]))
	}

	return score
}

func convertStringToNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return num
}
