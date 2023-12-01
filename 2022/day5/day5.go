package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	move int
	from int
	to   int
}

func getMovesList() []string {
	data, err := os.ReadFile("./inputs/day5.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	strData := string(data)
	idx := strings.Index(strData, "move")

	moves := strData[idx:]
	movesArr := strings.Split(moves, "\n")

	return movesArr
}

func Organize() []string {
	moves := getMovesList()

	stacks := []*Stack{
		{crates: []byte{'W', 'M', 'L', 'F'}, length: 4},
		{crates: []byte{'B', 'Z', 'V', 'M', 'F'}, length: 5},
		{crates: []byte{'H', 'V', 'R', 'S', 'L', 'Q'}, length: 6},
		{crates: []byte{'F', 'S', 'V', 'Q', 'P', 'M', 'T', 'J'}, length: 8},
		{crates: []byte{'L', 'S', 'W'}, length: 3},
		{crates: []byte{'F', 'V', 'P', 'M', 'R', 'J', 'W'}, length: 7},
		{crates: []byte{'J', 'Q', 'C', 'P', 'N', 'R', 'F'}, length: 7},
		{crates: []byte{'V', 'H', 'P', 'S', 'Z', 'W', 'R', 'B'}, length: 8},
		{crates: []byte{'B', 'M', 'J', 'C', 'G', 'H', 'Z', 'W'}, length: 8},
	}

	return moveCrates(moves, stacks)
}

func parseMove(moveStr string) Move {

	moves := strings.Index(moveStr, "move")
	from := strings.Index(moveStr, "from")
	to := strings.Index(moveStr, "to")

	moveAmt, err := strconv.Atoi(moveStr[moves+5 : from-1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	fromStack, err := strconv.Atoi(moveStr[from+5 : to-1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	toStack, err := strconv.Atoi(moveStr[to+3:])
	if err != nil {
		fmt.Println("Error:", err)
	}

	return Move{
		move: moveAmt,
		from: fromStack,
		to:   toStack,
	}
}

func moveCrates(list []string, stacks []*Stack) []string {
	var topCrates []string

	for _, move := range list {
		moveData := parseMove(move)
		fromStack := stacks[moveData.from-1]
		toStack := stacks[moveData.to-1]

		var temp []byte

		for i := 0; i < moveData.move; i++ {
			// crate := fromStack.pop()
			// toStack.insert(crate)
			temp = append(temp, fromStack.pop())
		}

		for i := len(temp) - 1; i >= 0; i-- {
			toStack.insert(temp[i])
		}

	}

	for _, stack := range stacks {
		topCrates = append(topCrates, string(stack.peek()))
	}

	return topCrates
}
