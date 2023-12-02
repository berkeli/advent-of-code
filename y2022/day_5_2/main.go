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

	chunks := strings.Split(string(f), "\n\n")

	stack, actions := chunks[0], chunks[1]

	boxes := make(map[string][]string)

	stacks := strings.Split(stack, "\n")

	reversed := reverse(stacks)

	keys := strings.Fields(reversed[0])

	for _, line := range reversed[1:] {
		for i, key := range keys {
			ind := 1 + 4*i
			if string(line[ind]) == " " {
				continue
			}
			boxes[key] = append(boxes[key], string(line[ind]))
		}
	}

	for _, action := range strings.Split(actions, "\n") {
		n, from, to := 0, "0", "0"

		fmt.Sscanf(action, "move %d from %s to %s", &n, &from, &to)

		boxes[to] = append(boxes[to], boxes[from][len(boxes[from])-n:]...)
		boxes[from] = boxes[from][:len(boxes[from])-n]
	}

	for _, key := range keys {
		fmt.Printf("%s", boxes[key][len(boxes[key])-1])
	}
}

func reverse(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
