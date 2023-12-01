package day14

import (
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

func getData() {
	data, err := os.ReadFile("./inputs/day14_test.txt")
	// data, err := os.ReadFile("./inputs/day14.txt")
	if err != nil {
		panic(err)
	}

	var parsedLines [][]image.Point

	strLines := strings.Split(string(data), "\n")
	for _, str := range strLines {
		var parsedCoords []image.Point

		coords := strings.Split(str, " -> ")
		for _, point := range coords {
			p := strings.Split(point, ",")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			parsedCoords = append(parsedCoords, image.Point{X: x, Y: y})
		}

		parsedLines = append(parsedLines, parsedCoords)
	}

	// fmt.Println(parsedLines)
	// origin := image.Point{500,0}
	originRow := []string{"."}
	matrix := [][]string{originRow}

	for _, lines := range parsedLines {
		for _, point := range lines {
			x := point.X

			diff := math.Abs(500.0 - float64(x))
			if x > 500 {
				for i := 0; i < int(diff); i++ {
					originRow = append([]string{"."}, originRow...)
				}
			} else {
				for i := 0; i < int(diff); i++ {
					originRow = append(originRow, ".")
				}
			}
		}
	}

	for _, row := range matrix {
		fmt.Println(row)
	}

}

func Day14() string {
	getData()
	return "..."
}
