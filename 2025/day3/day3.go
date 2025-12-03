package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFileIntoLines("2025/day3/input")
	type joltage struct {
		pos   []int
		value []int
	}
	total := 0
	jolt_size := 12
	for _, line := range lines {
		jolt := joltage{}
		jolt.pos = make([]int, 0)
		jolt.value = make([]int, 0)
		joltval := ""
		for js := 0; js < jolt_size; js++ {
		outer:
			for i := 9; i > 0; i-- {
				for i2, s := range line {
					if i2 > len(line)-(jolt_size-js) || (js > 0 && i2 <= jolt.pos[js-1]) {
						continue
					}
					if int(s-'0') == i {
						jolt.pos = append(jolt.pos, i2)
						jolt.value = append(jolt.value, int(s-'0'))
						joltval += string(s)
						break outer
					}
				}
			}
		}
		fmt.Println(line, joltval)
		total += utils.ToInt(joltval)
	}
	fmt.Println(total)
}
