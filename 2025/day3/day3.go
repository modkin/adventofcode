package main

import (
	"adventofcode/utils"
	"fmt"
)

func get_max_joltage(jolt_size int, lines []string) (total_joltage int) {

	for _, line := range lines {
		pos := make([]int, 0)
		joltval := ""
		for js := 0; js < jolt_size; js++ {
		outer:
			for i := 9; i > 0; i-- {
				for i2, s := range line {
					if i2 > len(line)-(jolt_size-js) || (js > 0 && i2 <= pos[js-1]) {
						continue
					}
					if int(s-'0') == i {
						pos = append(pos, i2)
						joltval += string(s)
						break outer
					}
				}
			}
		}
		total_joltage += utils.ToInt(joltval)
	}
	return
}

func main() {
	lines := utils.ReadFileIntoLines("2025/day3/input")

	fmt.Println("Day 3.1:", get_max_joltage(2, lines))
	fmt.Println("Day 3.2:", get_max_joltage(12, lines))
}
