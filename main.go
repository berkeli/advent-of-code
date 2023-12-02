package main

import (
	"log"
	"os"

	"github.com/berkeli/advent-of-code/y2023"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		panic("not enough arguments, please follow the format: YEAR DAY PART PATH, e.g. 2023 1 1 2023/1/input.txt\n PATH is optional")
	}

	switch args[0] {
	case "2023":
		printResult(y2023.Run(args[1:]))
	}
}

func printResult(res any, err error) {
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
