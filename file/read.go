package file

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ReadFromFile() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([][]string, 0, 0)

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Compile a regular expression to match one or more spaces
		re := regexp.MustCompile(`\s+`)

		// Replace multiple spaces with a single space
		line = re.ReplaceAllString(line, " ")
		data = append(data, []string{line})
	}

	return data
}
