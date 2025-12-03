package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadFileIntoLines("2025/day3/input")
	type joltage struct {
		pos   []int
		value []int
	}
	total := 0
	for _, line := range lines {
		jolt := joltage{}
	outer:
		for i := 9; i > 0; i-- {
			for i2, s := range line {
				if i2 == len(line)-1 {
					continue
				}
				if int(s-'0') == i {
					jolt.pos1 = i2
					jolt.value1 = int(s - '0')
					break outer
				}
			}
		}
	outer2:
		for i := 9; i > 0; i-- {
			for i2, s := range line {
				if i2 == 0 || i2 <= jolt.pos1 {
					continue
				}
				if int(s-'0') == i {
					jolt.pos2 = i2
					jolt.value2 = int(s - '0')
					break outer2
				}
			}
		}
		fmt.Println(jolt.value1, jolt.value2)
		total += utils.ToInt(strconv.Itoa(jolt.value1) + strconv.Itoa(jolt.value2))
	}
	fmt.Println(total)
}
