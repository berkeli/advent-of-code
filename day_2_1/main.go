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
	me = map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
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

		oppMove := opponent[s[0]]
		myMove := me[s[1]]

		r := checkScore(oppMove, myMove)

		if r == "win" {
			myScore += points[myMove] + 6
		}

		if r == "lose" {
			myScore += points[myMove]
		}

		if r == "draw" {
			myScore += points[myMove] + 3
		}
	}

	println(myScore)
}

func checkScore(opponentMove, myMove string) string {
	if opponentMove == myMove {
		return "draw"
	}

	if opponentMove == "rock" {
		if myMove == "paper" {
			return "win"
		}

		return "lose"
	}

	if opponentMove == "paper" {
		if myMove == "scissors" {
			return "win"
		}

		return "lose"
	}

	if opponentMove == "scissors" {
		if myMove == "rock" {
			return "win"
		}

		return "lose"
	}

	return "draw"
}
