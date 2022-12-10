package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CPU struct {
	x int
}

func main() {

	file, err := os.Open("2022/day10/input")
	if err != nil {
		panic(err)
	}

	//grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	crt := make(map[[2]int]string)
	cpu := CPU{1}
	instructions := make([][]string, 0)
	pipeline := make([][]string, 0)

	for scanner.Scan() {
		instructions = append(instructions, strings.Split(scanner.Text(), " "))

	}

	for _, instruction := range instructions {

		if instruction[0] == "noop" {
			pipeline = append(pipeline, instruction)
		} else if instruction[0] == "addx" {
			pipeline = append(pipeline, []string{"noop"})
			pipeline = append(pipeline, instruction)
		}
	}

	fmt.Println(pipeline)
	cycle := 1
	score := 0
	crtPos := 0
	for _, i2 := range pipeline {
		fmt.Println("B", cycle, cpu.x, score)
		pos := [2]int{crtPos % 40, crtPos / 40}
		if utils.IntAbs(cpu.x-crtPos%40) <= 1 {
			crt[pos] = "#"
		} else {
			crt[pos] = "."
		}

		if cycle == 20 {
			score += cpu.x * cycle
		} else if (cycle-20)%40 == 0 {
			score += cpu.x * cycle
		}
		if i2[0] == "addx" {
			cpu.x += utils.ToInt(i2[1])
		}
		fmt.Println("E", cycle, cpu.x, score)
		utils.Print2DStringsGrid(crt)
		cycle++
		crtPos++
	}

	fmt.Println("Day 10.2:", len(instructions))

}
