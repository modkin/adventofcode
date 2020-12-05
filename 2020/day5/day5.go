package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2020/day5/input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	row, column, maxId, minId := 0, 0, 0, math.MaxInt32
	seatList := make(map[int]bool)
	for scanner.Scan() {
		directions := strings.Split(scanner.Text(), "")
		for i := 0; i < 7; i++ {
			if directions[i] == "B" {
				row |= 1 << (6 - i)
			}
		}
		for i := 0; i < 3; i++ {
			if directions[7+i] == "R" {
				column |= 1 << (2 - i)
			}
		}
		id := row*8 + column
		seatList[id] = true
		if id > maxId {
			maxId = id
		}
		if id < minId {
			minId = id
		}
		row, column = 0, 0
	}

	fmt.Println("Task 5.1:", maxId)
	for i := minId; i <= maxId; i++ {
		if !seatList[i] {
			fmt.Println("Task 5.2:", i)
			break
		}
	}
}
