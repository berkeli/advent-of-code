package main

import (
	"os"
	"sort"
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

	groupSums := []int{}

	for _, group := range groups {
		calories := strings.Split(group, "\n")
		cur := 0
		for _, calorie := range calories {
			num, _ := strconv.Atoi(calorie)
			cur += num
		}

		groupSums = append(groupSums, cur)
	}

	sort.Ints(groupSums)

	println(groupSums[len(groupSums)-1] + groupSums[len(groupSums)-2] + groupSums[len(groupSums)-3])
}
