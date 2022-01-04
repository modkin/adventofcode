package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func print2D(input [][]string) {
	for _, i2 := range input {
		fmt.Println(i2)
	}
}

func move(floor [][]string) bool {
	moved := false
	floorCopy := make([][]string, len(floor))
	for i, _ := range floorCopy {
		floorCopy[i] = make([]string, len(floor[i]))
		copy(floorCopy[i], floor[i])
	}
	for i := 0; i < len(floor); i++ {
		for k := 0; k < len(floor[0]); k++ {
			if floorCopy[i][k] == ">" {
				targetIdx := (k + 1) % len(floor[0])
				if floorCopy[i][targetIdx] == "." {
					floor[i][k] = "."
					floor[i][targetIdx] = ">"
					moved = true
					k++
				}
			}
		}
	}
	for i, _ := range floorCopy {
		copy(floorCopy[i], floor[i])
	}
	for k := 0; k < len(floor[0]); k++ {
		for i := 0; i < len(floor); i++ {
			if floorCopy[i][k] == "v" {
				targetIdx := (i + 1) % len(floor)
				if floorCopy[targetIdx][k] == "." {
					floor[i][k] = "."
					floor[targetIdx][k] = "v"
					moved = true
					i++
				}
			}
		}
	}
	return moved
}

func main() {
	file, err := os.Open("2021/day25/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	floor := make([][]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		floor = append(floor, line)
	}
	//print2D(floor)
	//fmt.Println()
	//move(floor)
	//print2D(floor)
	//move(floor)
	//fmt.Println()
	//print2D(floor)
	counter := 1
	for move(floor) {
		counter++

	}
	//print2D(floor)
	fmt.Println(counter)
}
