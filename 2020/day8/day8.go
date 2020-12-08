package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

type instruction struct {
	inst string
	arg  int
}

func runProgramm(instList []instruction) (int, bool) {
	visited := make(map[int]bool)
	inf := false
	acc := 0
	pos := 0
	visited[0] = true
	for {
		current := instList[pos]
		if current.inst == "acc" {
			acc += current.arg
			pos++

		}
		if current.inst == "jmp" {
			pos += current.arg
		}
		if current.inst == "nop" {
			pos++
		}
		if visited[pos] {
			inf = true
			break
		} else if pos == len(instList) {
			break
		} else {
			visited[pos] = true
		}
	}
	return acc, inf
}

func main() {

	instList := make([]instruction, 0)

	scanner := bufio.NewScanner(utils.OpenFile("2020/day8/input"))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		instList = append(instList, instruction{line[0], utils.ToInt(line[1])})
	}
	//fmt.Println(instList)

	acc, _ := runProgramm(instList)
	fmt.Println("Task 8.1", acc)
	for i, elem := range instList {
		if elem.inst == "jmp" {
			instList[i] = instruction{"nop", elem.arg}
			acc, inf := runProgramm(instList)
			if !inf {
				fmt.Println("Task 8.2", acc)
				break
			} else {
				instList[i] = instruction{"jmp", elem.arg}
			}
		}
		if elem.inst == "nop" {
			instList[i] = instruction{"jmp", elem.arg}
			acc, inf := runProgramm(instList)
			if !inf {
				fmt.Println("Task 8.2", acc)
				break
			} else {
				instList[i] = instruction{"nop", elem.arg}
			}
		}
	}
}
