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
	visited := map[int]bool{0: true}
	inf := false
	acc, pos := 0, 0
	for {
		current := instList[pos]
		if current.inst == "acc" {
			acc += current.arg
			pos++
		} else if current.inst == "jmp" {
			pos += current.arg
		} else if current.inst == "nop" {
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

func change(input string) string {
	if input == "nop" {
		return "jmp"
	} else {
		return "nop"
	}
}

func main() {
	instList := make([]instruction, 0)
	scanner := bufio.NewScanner(utils.OpenFile("2020/day8/input"))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		instList = append(instList, instruction{line[0], utils.ToInt(line[1])})
	}

	acc, _ := runProgramm(instList)
	fmt.Println("Task 8.1", acc)

	for i, elem := range instList {
		if elem.inst == "jmp" || elem.inst == "nop" {
			instList[i] = instruction{change(elem.inst), elem.arg}
			acc, inf := runProgramm(instList)
			if !inf {
				fmt.Println("Task 8.2", acc)
				break
			} else {
				instList[i] = instruction{elem.inst, elem.arg}
			}
		}
	}
}
