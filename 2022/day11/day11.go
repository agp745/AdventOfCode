package day11

import (
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	startingItem []int
	operation    func(int) int
	testFunc     func(int) bool
	trueThrow    int
	falseThrow   int
}

func part2(input []byte) {
	var MonkeyList []Monkey
	Biggest := 1
	for _, monkey := range strings.Split(string(input), "\n\n") {
		m, b := parseMonkey(monkey)
		MonkeyList = append(MonkeyList, m)
		Biggest *= b
	}
	worryLevels := make([][]int, len(MonkeyList))
	for i, monkey := range MonkeyList {
		worryLevels[i] = monkey.startingItem
	}
	inspectCount := make([]int, len(MonkeyList))

	for i := 0; i < 10000; i++ {
		for j, monkey := range MonkeyList {
			for k, item := range worryLevels[j] {
				worryLevels[j][k] = monkey.operation(item)
				worryLevels[j][k] = worryLevels[j][k] % Biggest
				inspectCount[j]++
				if monkey.testFunc(worryLevels[j][k]) {
					worryLevels[monkey.trueThrow] = append(worryLevels[monkey.trueThrow], worryLevels[j][k])
				} else {
					worryLevels[monkey.falseThrow] = append(worryLevels[monkey.falseThrow], worryLevels[j][k])
				}
			}
			worryLevels[j] = []int{}
		}
	}

	max := 0
	secondMax := 0
	for _, count := range inspectCount {
		if count > max {
			secondMax = max
			max = count
		} else if count > secondMax {
			secondMax = count
		}
	}
	println("Part 2:", max*secondMax)
}

func part1(input []byte) {
	var MonkeyList []Monkey
	for _, monkey := range strings.Split(string(input), "\n\n") {
		m, _ := parseMonkey(monkey)
		MonkeyList = append(MonkeyList, m)
	}
	worryLevels := make([][]int, len(MonkeyList))
	for i, monkey := range MonkeyList {
		worryLevels[i] = monkey.startingItem
	}
	inspectCount := make([]int, len(MonkeyList))

	for i := 0; i < 20; i++ {
		for j, monkey := range MonkeyList {
			for k, item := range worryLevels[j] {
				worryLevels[j][k] = monkey.operation(item)
				worryLevels[j][k] /= 3
				inspectCount[j]++
				if monkey.testFunc(worryLevels[j][k]) {
					worryLevels[monkey.trueThrow] = append(worryLevels[monkey.trueThrow], worryLevels[j][k])
				} else {
					worryLevels[monkey.falseThrow] = append(worryLevels[monkey.falseThrow], worryLevels[j][k])
				}
			}
			worryLevels[j] = worryLevels[j][:0]
		}
	}

	// Find the two monkeys with the most inspections
	max := 0
	secondMax := 0
	for _, count := range inspectCount {
		if count > max {
			secondMax = max
			max = count
		} else if count > secondMax {
			secondMax = count
		}
	}
	println("Part 1:", max*secondMax)
}

func parseMonkey(monkey string) (Monkey, int) {
	var Biggest = 1
	m := Monkey{}
	for _, s := range strings.Split(monkey, "\n") {
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "Starting items:") {
			s = strings.Replace(s, "Starting items:", "", 1)
			s = strings.TrimSpace(s)
			for _, s2 := range strings.Split(s, ",") {
				s2 = strings.TrimSpace(s2)
				i, err := strconv.Atoi(s2)
				if err != nil {
					panic(err)
				}
				m.startingItem = append(m.startingItem, i)
			}
		} else if strings.HasPrefix(s, "Operation:") {
			s = strings.Replace(s, "Operation:", "", 1)
			s = strings.Replace(s, "new = old", "", 1)
			s = strings.TrimSpace(s)
			sign := s[0]
			s = s[1:]
			s = strings.TrimSpace(s)
			i, err := strconv.Atoi(s)
			if err != nil {
				m.operation = func(old int) int {
					return old * old
				}
			} else {
				if sign == '*' {
					m.operation = func(old int) int {
						return old * i
					}
				} else {
					m.operation = func(old int) int {
						return old + i
					}
				}
			}
		} else if strings.HasPrefix(s, "Test:") {
			s = strings.Replace(s, "Test:", "", 1)
			s = strings.Replace(s, "divisible by", "", 1)
			s = strings.TrimSpace(s)
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			m.testFunc = func(old int) bool {
				return old%i == 0
			}
			Biggest *= i
		} else if strings.HasPrefix(s, "If true:") {
			s = strings.Replace(s, "If true:", "", 1)
			s = strings.Replace(s, "throw to monkey", "", 1)
			s = strings.TrimSpace(s)
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			m.trueThrow = i
		} else if strings.HasPrefix(s, "If false:") {
			s = strings.Replace(s, "If false:", "", 1)
			s = strings.Replace(s, "throw to monkey", "", 1)
			s = strings.TrimSpace(s)
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			m.falseThrow = i
		}
	}
	return m, Biggest
}

func Day11() {
	input, err := os.ReadFile("./inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}
