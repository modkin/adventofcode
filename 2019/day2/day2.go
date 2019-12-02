package main

import (
	"adventofcode/2018/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

func compute(opcode int, first int, second int) int {
	if opcode == 1 {
		return first + second
	} else if opcode == 2 {
		return first * second
	}
	return -1
}

func processIntCode() int {
	content, err := ioutil.ReadFile("2019/day2/input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt(elem)
	}
	///adjust according to task
	intcode[1] = 12
	intcode[2] = 2

	index := 0
	for true {
		opCode := intcode[index]
		if opCode == 99 {
			fmt.Println(intcode)
			return intcode[0]
		}
		firstPos := intcode[index+1]
		secondPos := intcode[index+2]
		resultPos := intcode[index+3]
		intcode[resultPos] = compute(opCode, intcode[firstPos], intcode[secondPos])
		index += 4
	}
	return -1
}

func main() {
	fmt.Println("Task 2.1: ", processIntCode())
}
