package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	winning []int
	numbers []int
}

func part2(cards []card, idx int, copies int) {
	curr := cards[idx]
	matches := 0
	for _, w := range curr.winning {
		if slices.Contains(curr.numbers, w) {
			matches++
		}
	}
}

func main() {
	// data, _ := os.ReadFile("./inputs/day4_test.txt")
	data, _ := os.ReadFile("./inputs/day4.txt")

	re := regexp.MustCompile(`\d+`)
	cards := []card{}

	for _, s := range strings.Split(string(data), "\n") {
		var c card
		nums := strings.Split(s, ":")
		nums = strings.Split(strings.TrimSpace(nums[1]), "|")

		for _, m := range re.FindAllStringSubmatch(nums[0], -1) {
			n, _ := strconv.Atoi(m[0])
			c.winning = append(c.winning, n)
		}
		for _, m := range re.FindAllStringSubmatch(nums[1], -1) {
			n, _ := strconv.Atoi(m[0])
			c.numbers = append(c.numbers, n)
		}
		cards = append(cards, c)
	}

	p1 := 0
	for _, card := range cards {
		matches := 0
		for _, w := range card.winning {
			if slices.Contains(card.numbers, w) {
				matches++
			}
		}
		points := math.Pow(float64(2), float64(matches-1))
		p1 += int(points)
	}

	p2 := 0
	for _, card := range cards {
		matches := 0
		for _, w := range card.winning {
			if slices.Contains(card.numbers, w) {
				matches++
			}
		}

	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
}
