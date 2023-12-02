package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	content := string(f)

	groups := strings.Split(content, "\n\n")

	maxSum := 0

	for _, group := range groups {
		calories := strings.Split(group, "\n")
		cur := 0
		for _, calorie := range calories {
			num, _ := strconv.Atoi(calorie)
			cur += num
		}

		if cur > maxSum {
			maxSum = cur
		}
	}

	println(maxSum)
}
