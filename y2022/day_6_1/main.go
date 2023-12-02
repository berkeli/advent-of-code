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

	input := string(f)

	curr := ""

	for i, ch := range input {
		if len(curr) == 3 && !contains(curr, ch) {
			fmt.Println(i + 1)
			return
		}

		if !contains(curr, ch) {
			curr += string(ch)
		} else {
			j := strings.IndexRune(curr, ch)
			curr = input[i-len(curr)+j+1 : i+1]
		}
	}
}

func contains(a string, b rune) bool {
	for _, c := range a {
		if c == b {
			return true
		}
	}
	return false
}
