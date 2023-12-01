package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name   string
	amount int
}

const (
	ADDX = "addx"
	NOOP = "noop"
)

func getData() []Instruction {
	data, err := os.ReadFile("./inputs/day10.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	dataArr := strings.Split(string(data), "\n")
	var instructions []Instruction

	for _, str := range dataArr {
		arr := strings.Split(str, " ")
		if len(arr) == 2 {
			n, err := strconv.Atoi(arr[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
			instructions = append(instructions, Instruction{arr[0], n})
		} else {
			instructions = append(instructions, Instruction{arr[0], 0})
		}
	}

	return instructions
}

func Day10() int {
	data := getData()
	cpu := InitCPU()

	parseInstructions(cpu, data)

	for _, arr := range cpu.CRT.Matrix {
		fmt.Println(arr)
	}
	return cpu.CombinedSignals
}

func parseInstructions(cpu *CPU, data []Instruction) {
	for _, instruction := range data {
		if instruction.name == ADDX {
			cpu.addX(instruction.amount)
		} else {
			cpu.noop()
		}
	}
}
