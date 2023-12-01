package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PROMPT = "$"
	LS     = "ls"
	CD     = "cd"
	BACK   = ".."

	MAX = 100000
)

func getData() []string {
	data, err := os.ReadFile("./inputs/day7.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	dataArr := strings.Split(string(data), "\n")
	return dataArr
}

func Day7() int {
	commands := getData()
	fs := NewFS()

	//build filesystem
	for i, cmd := range commands {
		tokens := strings.Split(cmd, " ")

		if tokens[1] == CD {
			cdCmd(fs, tokens[2])
		}

		if tokens[1] == LS {
			lsCmd(fs, commands, i+1)
		}
	}

	//Finish fs by calculating correct dir sizes
	var totalSize int
	traverseFS(fs, fs.root, &totalSize)

	availSpace := 70000000 - fs.root.size
	minSpace := 30000000 - availSpace
	// fmt.Println("AVAILABLE", availSpace)
	fmt.Println("MIN", minSpace)
	size := findMinSpace(fs, minSpace)

	return size
}

func traverseFS(fs *FileSystem, curr *dir, total *int) {
	if curr == nil {
		return
	}

	for _, d := range curr.dirs {
		traverseFS(fs, d, total)
		curr.size += d.size
	}

	if curr.size > MAX {
		return
	}

	*total += curr.size
}

func findMinSpace(fs *FileSystem, minSpace int) int {
	var dirSizes []int
	q := []*dir{fs.root}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.size >= minSpace {
			dirSizes = append(dirSizes, curr.size)
		}

		for _, d := range curr.dirs {
			q = append(q, d)
		}
	}

	min := dirSizes[0]
	for _, size := range dirSizes {
		if size < min {
			min = size
		}
	}

	return min
}

func cdCmd(fs *FileSystem, arg string) {
	if arg == BACK {
		fs.CDBack()
	} else {
		fs.CDInto(arg)
	}
}

func lsCmd(fs *FileSystem, commands []string, idx int) {
	for i := idx; i < len(commands); i++ {
		tokens := strings.Split(commands[i], " ")
		if tokens[0] == PROMPT {
			break
		}
		if tokens[0] == "dir" {
			continue
		}

		fileSize, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Println("Error: could not convert string to int")
		}
		fs.GetFile(tokens[1], fileSize)
	}
}
