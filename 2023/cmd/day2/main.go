package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./inputs/day2.txt")
	re := regexp.MustCompile(`(\d+) (\w+)`)

	total := 0
	power := 0

	for i, s := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		min := map[string]int{}

		for _, m := range re.FindAllStringSubmatch(s, -1) {
			n, _ := strconv.Atoi(m[1])

			min[m[2]] = slices.Max([]int{min[m[2]], n})
		}

		//part 1
		if min["red"] <= 12 && min["green"] <= 13 && min["blue"] <= 14 {
			total += i + 1
		}

		//part 2
		p := min["red"] * min["green"] * min["blue"]
		power += p
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", total, power)
}
