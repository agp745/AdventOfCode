package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Overlaps() int {

	data, err := os.ReadFile("./inputs/day4.txt")
	if err != nil {
		fmt.Println("ERROR READING FILE:\n", err)
	}

	pairs := strings.Split(string(data), "\n")

	return countOverlaps(pairs)
}

func countOverlaps(arr []string) int {
	var overlaps int

	for _, pair := range arr {
		temp := strings.Split(pair, ",")

		a := strings.Split(temp[0], "-")
		b := strings.Split(temp[1], "-")

		x1, y1 := convertToInt(a)
		x2, y2 := convertToInt(b)

		//part 1
		// if x1 >= x2 && y1 <= y2 {
		// 	overlaps++
		// } else if x2 >= x1 && y2 <= y1 {
		// 	overlaps++
		// }

		if isOverlapped(x1, y1, x2, y2) {
			overlaps++
		}
	}

	return overlaps
}

func convertToInt(arr []string) (int, int) {
	n1, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	n2, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	return n1, n2
}

func isOverlapped(x1 int, y1 int, x2 int, y2 int) bool {
	if x1 >= x2 && x1 <= y2 {
		return true
	}

	if x2 >= x1 && x2 <= y1 {
		return true
	}

	if y1 >= x2 && y1 <= y2 {
		return true
	}

	if y2 >= x1 && y2 <= y1 {
		return true
	}

	return false
}
