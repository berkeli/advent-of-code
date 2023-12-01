package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int64
}

type Folder struct {
	Parent string
	Files  []File
	Size   int64
}

func main() {
	f, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	input := string(f)

	folders := make(map[string]Folder)

	var currDir string

	arr := strings.Split(input, "\n")
	i := 0
	for i < len(arr) {
		line := strings.Split(arr[i], " ")
		fmt.Println(line)
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == ".." {
					currDir = folders[currDir].Parent
					i++
					continue
				}
				dirName := ""
				if currDir == "" {
					dirName = line[2]
				} else {
					dirName = fmt.Sprintf("%s->%s", currDir, line[2])
				}
				folders[dirName] = Folder{Parent: currDir}
				currDir = dirName
			}
			i++
		} else {
			for i < len(arr) && !strings.HasPrefix(arr[i], "$") {
				line := strings.Split(arr[i], " ")

				if line[0] == "dir" {
					folders[currDir+"->"+line[1]] = Folder{Parent: currDir}
					i++
					continue
				}
				dir := folders[currDir]
				size, _ := strconv.ParseInt(line[0], 10, 64)
				dir.Files = append(folders[currDir].Files, File{Name: line[1], Size: size})
				folders[currDir] = dir
				i++
			}
		}

	}

	sum := int64(0)

	for k, v := range folders {
		if len(v.Files) > 0 {
			for _, f := range v.Files {
				v.Size += f.Size
			}
			folders[k] = v
		}
		if folders[k].Parent != "" {
			f := folders[folders[k].Parent]
			f.Size += folders[k].Size
			folders[folders[k].Parent] = f
		}
	}

	for _, v := range folders {
		if v.Size < 100000 {
			sum += v.Size
		}
	}

	fmt.Println(sum)
}
