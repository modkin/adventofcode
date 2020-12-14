package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"strings"
)

func applyMask(value int, mask []int) int {
	//fmt.Printf("%036b\n", value)
	for i, elem := range mask {
		if elem == 0 {
			value = value & ((math.MaxInt64) ^ 1<<(35-i))
		} else if elem == 1 {
			value = value | 1<<(35-i)
		}
	}
	return value
}

func applyMemoryMask(address int, mask []int) []int {
	ret := make([]int, 0)
	for i, elem := range mask {
		if elem == 1 {
			address = address | 1<<(35-i)
		}
	}
	var permutate func(int, int)
	permutate = func(address int, i int) {
		if i == 36 {
			return
		}
		if mask[i] == 4 {
			ret = append(ret, address)
			addressTmp := address ^ 1<<(35-i)
			ret = append(ret, addressTmp)
			permutate(addressTmp, i+1)
		}
		permutate(address, i+1)
	}
	permutate(address, 0)
	return ret
}

func main() {

	memory := make(map[int]int)
	memory2 := make(map[int]int)
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
			multipleAdr := applyMemoryMask(utils.ToInt(address), mask[0:])
			for _, adr := range multipleAdr {
				memory2[adr] = utils.ToInt(line[1])
			}
		}
	}
	sum := 0
	for _, elem := range memory {
		sum += elem
	}
	fmt.Println("Task 14.1:", sum)
	sum2 := 0
	for _, elem := range memory2 {
		sum2 += elem
	}
	fmt.Println("Task 14.2:", sum2)
}
