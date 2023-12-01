package day6

import (
	"fmt"
	"os"
)

func getData() string {
	data, err := os.ReadFile("./inputs/day6.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return string(data)
}

func FindMarker() int {
	data := getData()
	markerLen := 14

	for i := 0; i < len(data)-markerLen; i++ {
		sub := data[i : i+markerLen]
		unique := make(map[byte]bool)

		for _, char := range sub {
			if unique[byte(char)] {
				break
			}
			unique[byte(char)] = true
		}

		if len(unique) == markerLen {
			return i + markerLen
		}
	}

	return -1
}
