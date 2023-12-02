package d1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stringNum = []string{"oneight", "twone", "eightwo", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var num = []string{"18", "21", "82", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Run(part, path string) (int, error) {
	if path == "" {
		path = "input.txt"
	}
	f, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("could not read file: %w", err)
	}

	switch part {
	case "1":
		return partOne(f)
	case "2":
		return partTwo(f)
	default:
		return 0, fmt.Errorf("invalid part %v", part)
	}

}

func partOne(f *os.File) (int, error) {
	scanner := bufio.NewScanner(f)

	var sum int

	for scanner.Scan() {
		sum += extractNumber(scanner.Text())
	}

	return sum, nil
}

func partTwo(f *os.File) (int, error) {
	scanner := bufio.NewScanner(f)
	var sum int

	for scanner.Scan() {
		s := scanner.Text()
		for i, n := range stringNum {
			s = strings.ReplaceAll(s, n, num[i])
		}
		i := extractNumber(s)
		sum += i
	}

	return sum, nil
}

func extractNumber(s string) int {
	curr := []rune{}
	for _, r := range s {
		if r < '0' || r > '9' {
			continue
		}

		curr = append(curr, r)
	}

	return int(curr[0]-'0')*10 + int(curr[len(curr)-1]-'0')
}
