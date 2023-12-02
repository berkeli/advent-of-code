package y2023

import (
	"fmt"

	"github.com/berkeli/advent-of-code/y2023/d1"
	"github.com/berkeli/advent-of-code/y2023/d2"
)

const YEAR = "2023"

func Run(args []string) (int, error) {
	path := getPath(args)

	switch args[0] {
	case "1":
		res, err := d1.Run(args[1], path)
		if err != nil {
			return 0, err
		}

		return res, nil
	case "2":
		res, err := d2.Run(args[1], path)
		if err != nil {
			return 0, err
		}

		return res, nil
	default:
		return 0, fmt.Errorf("invalid day: %v", args[0])
	}
}

func getPath(args []string) string {
	// if path was not provided
	if len(args) < 3 {
		return fmt.Sprintf("y%v/d%v/input.txt", YEAR, args[0])
	}

	return args[2]
}
