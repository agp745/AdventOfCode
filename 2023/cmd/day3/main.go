package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func directions() []image.Point {
	return []image.Point{
		{0, -1},  // up
		{1, -1},  // up-right
		{1, 0},   // right
		{1, 1},   // down-right
		{0, 1},   // down
		{-1, 1},  // down-left
		{-1, 0},  // left
		{-1, -1}, // up-left
	}
}

func main() {
	input, _ := os.ReadFile("./inputs/day3.txt")

	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				grid[image.Point{x, y}] = r
			}
		}
	}

	p1, p2 := 0, 0
	parts := map[image.Point][]int{}
	for y, s := range strings.Fields(string(input)) {
		for _, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(s, -1) {
			bounds := map[image.Point]struct{}{}
			for x := m[0]; x < m[1]; x++ {
				for _, d := range directions() {
					bounds[image.Point{x, y}.Add(d)] = struct{}{}
				}
			}

			n, _ := strconv.Atoi(s[m[0]:m[1]])
			for p := range bounds {
				if _, ok := grid[p]; ok {
					parts[p] = append(parts[p], n)
					p1 += n
				}
			}
		}
	}

	for p, ns := range parts {
		if grid[p] == '*' && len(ns) == 2 {
			p2 += ns[0] * ns[1]
		}
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
}
