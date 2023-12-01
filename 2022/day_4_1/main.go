package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	count := 0

	lines := strings.Split(string(f), "\n")

	for _, line := range lines {
		pair1, pair2 := strings.Split(line, ",")[0], strings.Split(line, ",")[1]

		x1, y1 := split(pair1)
		x2, y2 := split(pair2)

		if (x1 >= x2 && y1 <= y2) || (x1 <= x2 && y1 >= y2) {
			count++
		}

	}

	fmt.Println(count)
}

func split(s string) (int, int) {
	a := strings.Split(s, "-")

	x, _ := strconv.Atoi(a[0])
	y, _ := strconv.Atoi(a[1])

	return x, y
}
