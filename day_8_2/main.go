package main

import (
	"fmt"
	"os"
	"strings"
)

func look(direction []int, treeHeight int, invert bool) bool {
	d := make([]int, len(direction))
	copy(d, direction)
	if invert {
		for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
			d[i], d[j] = d[j], d[i]
		}
	}
	for _, v := range d {
		if treeHeight <= v {
			return false
		}
	}
	return true
}

func treeInfo(i int, j int, grid [][]int) bool {
	h := grid[i][j]
	col := make([]int, len(grid[0]))
	for c := range grid {
		col[c] = grid[c][j]
	}
	l := look(grid[i][:j], h, true)    // left
	r := look(grid[i][j+1:], h, false) // right
	t := look(col[:i], h, true)        // top
	b := look(col[i+1:], h, false)     // bottom
	return l || r || t || b
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, treeHeight := range line {
			grid[i][j] = int(treeHeight - '0')
		}
	}
	partOne := (len(grid)*2 + len(grid[0])*2) - 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			visible := treeInfo(i, j, grid)
			if visible {
				partOne++
			}

		}
	}
	fmt.Println(partOne)
}
