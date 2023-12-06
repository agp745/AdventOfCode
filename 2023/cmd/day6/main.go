package main

import "fmt"

func getDistance(time int, speed int) int {
	return time * speed
}

func validTimes(time int, record int) int {
	total := 0
	for i := 1; i < time; i++ {
		remainingTime := time - i
		if getDistance(remainingTime, i) > record {
			total += 1
		}
	}
	return total
}

func main() {
	p1 := 1
	data := map[int]int{
		42686985: 284100511221341,
	}

	for k, v := range data {
		n := validTimes(k, v)
		p1 *= n
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, -1)
}
