package main

import (
	"os"

	dayone "github.com/berkeli/advent-of-code/2023/1"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("no test selected")
	}

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
			res, err := dayone.Run(args[2])
			if err != nil {
				panic(err)
			}

			println(res)
		}
	}

}
