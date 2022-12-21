package main

import (
	"os"
	"strings"
)

var (
	opponent = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}
	points = map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
)

func main() {
	f, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	content := string(f)

	games := strings.Split(content, "\n")
	myScore := 0

	for _, game := range games {
		s := strings.Split(game, " ")

		opp, res := s[0], s[1]

		switch res {
		case "X": // lose
			myScore += lose(opponent[opp])
		case "Y": // draw
			myScore += points[opponent[opp]] + 3
		case "Z": // win
			myScore += win(opponent[opp]) + 6
		}
	}

	println(myScore)
}

func lose(oppMove string) int {
	switch oppMove {
	case "rock":
		return points["scissors"]
	case "paper":
		return points["rock"]
	case "scissors":
		return points["paper"]
	}

	return 0
}

func win(oppMove string) int {
	switch oppMove {
	case "rock":
		return points["paper"]
	case "paper":
		return points["scissors"]
	case "scissors":
		return points["rock"]
	}

	return 0
}
