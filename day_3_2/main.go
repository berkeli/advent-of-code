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

	for _, line := range lines {
		left, right := line[0:len(line)/2], line[len(line)/2:]
		hash := make(map[string]bool)
		for i := 0; i < len(left); i++ {
			if contains(right, string(left[i])) && !hash[string(left[i])] {
				hash[string(left[i])] = true
				sum += priorities[string(left[i])]
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
