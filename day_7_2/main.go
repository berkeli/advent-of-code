package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	Parent *Folder
	Childs map[string]*Folder
	Size   int64
}

var sizes []int64

func (f *Folder) hydrateSize() {
	for _, v := range f.Childs {
		v.hydrateSize()
		sizes = append(sizes, v.Size)
		f.Size += v.Size
	}
}

func main() {
	f, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	input := string(f)

	root := Folder{Parent: nil, Childs: map[string]*Folder{}, Size: 0}

	var curr *Folder

	arr := strings.Split(input, "\n")
	i := 0
	for i < len(arr) {
		line := strings.Split(arr[i], " ")
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == "/" {
					curr = &root
					i++
					continue
				}
				if line[2] == ".." {
					curr = curr.Parent
					i++
					continue
				}
				curr = curr.Childs[line[2]]
			}
			i++
		} else {
			for i < len(arr) && !strings.HasPrefix(arr[i], "$") {
				line := strings.Split(arr[i], " ")

				if line[0] == "dir" {
					curr.Childs[line[1]] = &Folder{
						Parent: curr,
						Childs: map[string]*Folder{},
						Size:   0,
					}
					i++
					continue
				}

				size, _ := strconv.ParseInt(line[0], 10, 64)
				curr.Size += size
				i++
			}
		}

	}

	root.hydrateSize()

	spaceNeeded := -1 * (70000000 - root.Size - 30000000)

	sm := root.Size

	for _, v := range sizes {
		if v >= spaceNeeded && v < sm {
			sm = v
		}
	}

	fmt.Println(sm)
}
