package day8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func getData() [][]int {
	// data, err := os.ReadFile("./inputs/day8_test.txt")
	data, err := os.ReadFile("./inputs/day8.txt")
	if err != nil {
		fmt.Println("Error: could not read file")
	}

	dataArr := strings.Split(string(data), "\n")
	var strMatrix [][]string
	for _, str := range dataArr {
		strMatrix = append(strMatrix, strings.Split(str, ""))
	}

	var dataMatrix [][]int
	for _, arr := range strMatrix {
		var intArr []int
		for _, n := range arr {
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Printf("Error: could not convert %s to int", n)
			}
			intArr = append(intArr, num)
		}
		dataMatrix = append(dataMatrix, intArr)
	}

	return dataMatrix
}

func Day8() int {
	matrix := getData()
	totalVisible := (len(matrix) * 4) - 4

	start := Point{
		x: 1,
		y: 1,
	}

	seen := make([][]bool, len(matrix))
	for i := range seen {
		seen[i] = make([]bool, len(matrix[0]))
		for j := range seen[i] {
			seen[i][j] = false
		}
	}

	for i := range seen[0] {
		seen[0][i] = true
	}
	for i := range seen[len(seen)-1] {
		seen[len(seen[0])-1][i] = true
	}
	for i, arr := range seen {
		seen[i][0] = true
		seen[i][len(arr)-1] = true
	}

	bestScenery := 0

	Walk(matrix, start, seen, &totalVisible, &bestScenery)

	// return totalVisible
	return bestScenery
}
