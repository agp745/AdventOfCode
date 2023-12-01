package day13

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getData() []string {
	// data, err := os.ReadFile("./inputs/day13_test.txt")
	data, err := os.ReadFile("./inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(data))
}

func unmarshall(left, right string) (any, any) {
	var l, r any
	json.Unmarshal([]byte(left), &l)
	json.Unmarshal([]byte(right), &r)

	return l, r
}

func compare(left, right any) int {
	l, lok := left.([]any)
	r, rok := right.([]any)

	switch {
	case !lok && !rok:
		return int(left.(float64) - right.(float64))
	case !lok:
		l = []any{left}
	case !rok:
		r = []any{right}
	}

	for i := 0; i < len(l) && i < len(r); i++ {
		if res := compare(l[i], r[i]); res != 0 {
			return res
		}
	}

	return len(l) - len(r)
}

func isCorrectOrder(left, right string) bool {
	l, r := unmarshall(left, right)
	return compare(l, r) <= 0
}

func part1(data []string) int {
	sum := 0
	for i := 0; i < len(data); i += 2 {
		if isCorrectOrder(data[i], data[i+1]) {
			sum += (i/2 + 1)
		}
	}

	return sum
}

func part2(data []string) int {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {

			if ok := isCorrectOrder(data[j], data[j+1]); !ok {
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}

	result := 1
	for i, n := range data {
		if n == "[[2]]" || n == "[[6]]" {
			result *= (i + 1)
		}
	}

	return result
}

func Day13() string {
	data := getData()

	p1 := part1(data)
	p2 := part2(data)

	fmt.Printf("part1 => %v\n", p1)
	fmt.Printf("part2 => %v\n", p2)

	return "..."
}

// [[2]]
// [[6]]
