package d2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxAllowed = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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
		return 0, fmt.Errorf("invalid part: %v", part)
	}
}

func partOne(f *os.File) (int, error) {
	scanner := bufio.NewScanner(f)

	var sum int

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ":")

		if !isGameValid(text[1]) {
			continue
		}

		i, err := strconv.Atoi(strings.Split(text[0], " ")[1])
		if err != nil {
			return 0, err
		}

		sum += i
	}

	return sum, nil
}

func partTwo(f *os.File) (int, error) {
	scanner := bufio.NewScanner(f)
	var sum int

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ":")

		min := minCubesRequired(text[1])

		sum += min["red"] * min["green"] * min["blue"]
	}

	return sum, nil
}

func isGameValid(s string) bool {
	subsets := strings.Split(s, ";")
	for _, subset := range subsets {
		draws := strings.Split(subset, ",")
		for _, cubes := range draws {
			cube := strings.Split(strings.TrimSpace(cubes), " ")
			count, color := cube[0], cube[1]

			countInt, err := strconv.Atoi(count)
			if err != nil {
				log.Printf("invalid cube count: %v", err)
				return false
			}

			if countInt > maxAllowed[color] {
				return false
			}
		}
	}

	return true
}

func minCubesRequired(s string) map[string]int {
	min := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	subsets := strings.Split(s, ";")
	for _, subset := range subsets {
		draws := strings.Split(subset, ",")
		for _, cubes := range draws {
			cube := strings.Split(strings.TrimSpace(cubes), " ")
			count, color := cube[0], cube[1]

			countInt, err := strconv.Atoi(count)
			if err != nil {
				log.Printf("invalid cube count: %v", err)
				return nil
			}

			if countInt > min[color] {
				min[color] = countInt
			}
		}
	}

	return min
}
