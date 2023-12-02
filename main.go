package main

import (
	"fmt"
	"os"

	dayone "github.com/berkeli/advent-of-code/2023/1"
	daytwo "github.com/berkeli/advent-of-code/2023/2"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		panic("not enough arguments, please follow the format: YEAR DAY PART PATH, e.g. 2023 1 1")
	}

	path := getPath(args)

	switch args[0] {
	case "2023":
		switch args[1] {
		case "1":
			res, err := dayone.Run(args[2])
			if err != nil {
				panic(err)
			}

			println(res)
		case "2":
			res, err := daytwo.Run(args[2], path)
			if err != nil {
				panic(err)
			}

			println(res)
		}
	}

}

func getPath(args []string) string {
	if len(args) < 4 {
		return fmt.Sprintf("%v/%v/input.txt", args[0], args[1])
	}

	return args[3]
}
