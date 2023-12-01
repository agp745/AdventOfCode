package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	data, _ := os.ReadFile("./inputs/day1.txt")
	result := 0

	for _, s := range strings.Fields(string(data)) {
		var n1 string
		var n2 string
		for _, ch := range s {
			_, err := strconv.Atoi(string(ch))
			if err == nil {
				n1 = string(ch)
				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			ch := string(s[i])
			_, err := strconv.Atoi(ch)
			if err == nil {
				n2 = ch
				break
			}
		}

		n, _ := strconv.Atoi(n1 + n2)
		result += n
	}

	return result
}

func part2() int {
	// data, _ := os.ReadFile("./inputs/day1_test.txt")
	data, _ := os.ReadFile("./inputs/day1.txt")

	result := 0
	nums := map[string]string{
		"one":       "1",
		"two":       "2",
		"three":     "3",
		"four":      "4",
		"five":      "5",
		"six":       "6",
		"seven":     "7",
		"eight":     "8",
		"nine":      "9",
		"ten":       "0",
		"eleven":    "1",
		"twelve":    "2",
		"thirteen":  "3",
		"fourteen":  "4",
		"fifteen":   "5",
		"sixteen":   "6",
		"seventeen": "7",
		"eighteen":  "8",
		"nineteen":  "9",
		"1":         "1",
		"2":         "2",
		"3":         "3",
		"4":         "4",
		"5":         "5",
		"6":         "6",
		"7":         "7",
		"8":         "8",
		"9":         "9",
	}

	for _, s := range strings.Fields(string(data)) {
		var n1 string
		var n2 string

	outerloop1:
		for i := 0; i < len(s); i++ {
			for j := len(s); j >= i; j-- {
				val, ok := nums[s[i:j]]
				if ok {
					n1 = val
					break outerloop1
				}
			}
		}

	outerloop2:
		for i := len(s); i >= 0; i-- {
			for j := 0; j <= i; j++ {
				val, ok := nums[s[j:i]]
				if ok {
					n2 = val
					break outerloop2
				}
			}
		}

		n, _ := strconv.Atoi(n1 + n2)
		result += n
	}

	return result
}

func main() {
	p1 := part1()
	p2 := part2()
	fmt.Printf("part 1: %d\npart 2: %d\n", p1, p2)
}
