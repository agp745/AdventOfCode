package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Amount    int
}

func getData() []Instruction {
	// data, err := os.ReadFile("./inputs/day9_test2.txt")
	data, err := os.ReadFile("./inputs/day9_test.txt")
	// data, err := os.ReadFile("./inputs/day9.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	dataArr := strings.Split(string(data), "\n")
	instructions := []Instruction{}

	for _, str := range dataArr {
		arr := strings.Split(str, " ")
		dir := arr[0]
		amt, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println("Error:", err)
		}

		inst := Instruction{
			Direction: dir,
			Amount:    amt,
		}

		instructions = append(instructions, inst)
	}

	return instructions
}

func Day9() int {
	data := getData()
	rope := NewRope()

	parseInstructions(data, rope)

	Part2()

	return len(rope.visited)
}

func parseInstructions(instructions []Instruction, r *LinkedRope) {
	for _, i := range instructions {
		switch d := i.Direction; {
		case d == "L":
			r.MoveLeft(i.Amount)
		case d == "R":
			r.MoveRight(i.Amount)
		case d == "U":
			r.MoveUp(i.Amount)
		case d == "D":
			r.MoveDown(i.Amount)
		}
	}
}
