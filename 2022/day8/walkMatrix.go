package day8

func directions() []Point {
	directions := []Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	return directions
}

func checkVisibility(matrix [][]int, curr Point) bool {
	currValue := matrix[curr.y][curr.x]

	visible := false
	for _, dir := range directions() {
		x, y := curr.x, curr.y
		for {
			x += dir.x
			y += dir.y

			if x < 0 || y < 0 || x >= len(matrix[0]) || y >= len(matrix) {
				visible = true
				break
			}

			if matrix[y][x] >= currValue {
				break
			}
		}
	}

	return visible
}

func getSceneryValue(matrix [][]int, curr Point) int {
	currValue := matrix[curr.y][curr.x]
	var values []int

	for _, dir := range directions() {
		x, y := curr.x, curr.y
		var sceneryValue int

		for {
			x += dir.x
			y += dir.y

			if x < 0 || y < 0 || x >= len(matrix[0]) || y >= len(matrix) {
				break
			}

			if matrix[y][x] >= currValue {
				sceneryValue++
				break
			}

			sceneryValue++
		}
		values = append(values, sceneryValue)
	}
	totalScore := values[0] * values[1] * values[2] * values[3]
	return totalScore
}

func Walk(matrix [][]int, curr Point, seen [][]bool, visibleCount *int, bestScenery *int) {
	if curr.x < 1 || curr.y < 1 || curr.x >= len(matrix[0])-1 || curr.y >= len(matrix)-1 {
		return
	}

	if seen[curr.y][curr.x] {
		return
	}

	seen[curr.y][curr.x] = true

	if checkVisibility(matrix, curr) {
		*visibleCount++
	}

	sceneryValue := getSceneryValue(matrix, curr)

	if sceneryValue > *bestScenery {
		*bestScenery = sceneryValue
	}

	for _, dir := range directions() {
		x, y := curr.x+dir.x, curr.y+dir.y
		Walk(matrix, Point{x, y}, seen, visibleCount, bestScenery)
	}

}
