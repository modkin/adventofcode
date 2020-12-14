package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"strings"
)

func applyMask(value int, mask []int) int {
	//fmt.Println(value)
	//fmt.Printf("%036b\n", value)
	for i, elem := range mask {
		if elem == 0 {
			//fmt.Printf("%036b\n", (math.MaxInt64) ^ 1  << (35-i))
			value = value & ((math.MaxInt64) ^ 1<<(35-i))
			//fmt.Printf("%036b\n", value)
		} else if elem == 1 {
			value = value | 1<<(35-i)
			//fmt.Printf("%036b\n", value)
		}
	}
	//fmt.Println(value)
	return value
}

func main() {

	memory := make(map[int]int)
	var mask [36]int

	scanner := bufio.NewScanner(utils.OpenFile("2020/day14/input"))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")
		if line[0] == "mask" {
			for i, bit := range strings.Split(line[1], "") {
				if bit == "1" {
					mask[i] = 1
				} else if bit == "0" {
					mask[i] = 0
				} else if bit == "X" {
					mask[i] = 4
				}
			}
		} else {
			value := utils.ToInt(line[1])
			value = applyMask(value, mask[0:])
			address := strings.TrimSuffix(strings.Split(line[0], "mem[")[1], "]")
			memory[utils.ToInt(address)] = value
		}
	}
	sum := 0
	for _, elem := range memory {
		sum += elem
	}
	fmt.Println(sum)
}
