package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	priorities := make(map[string]int)
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		priorities[string(r)] = i
		i++
	}

	for r := 'A'; r <= 'Z'; r++ {
		priorities[string(r)] = i
		i++
	}

	sum := 0

	lines := strings.Split(string(f), "\n")

	for i := 0; i < len(lines); i += 3 {
		for j := 0; j < len(lines[i]); j++ {
			if contains(lines[i+1], string(lines[i][j])) && contains(lines[i+2], string(lines[i][j])) {
				sum += priorities[string(lines[i][j])]
				break
			}
		}
	}

	fmt.Println(sum)
}

func contains(s string, c string) bool {
	for _, r := range s {
		if string(r) == c {
			return true
		}
	}
	return false
}
