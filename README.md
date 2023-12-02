# My solutions to Advent of Code in GO

Years and days are packages, each day has two parts, each part is a function in the package. There will be tests based on examples provided on https://adventofcode.com/
I will try to backfill the previous years as I go.

## Usage

The main.go file is a wrapper for the solutions, it takes 4 arguments:
- YEAR - year of the challenge
- DAY - day of the challenge
- PART - part of the challenge
- INPUT_FILE_PATH - path to the input file. This is optional, if not provided, will use the input.txt file in the same directory as challenge day.

```bash
go run . YEAR DAY PART INTPU_FILE_PATH
```


## Example
```bash
go run . 2023 1 1 2023/1/input.txt
```

same as
```
go run . 2023 1 1
```

## Test

```bash
go test ./...
```
