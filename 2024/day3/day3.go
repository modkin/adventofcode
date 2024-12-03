package main

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"regexp"
)

func findUncorupptedWithDo(input string, part2 bool) (result int) {
	reg := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	if part2 {
		reg = regexp.MustCompile(`mul\((\d*),(\d*)\)|do\(\)|don't\(\)`)
	}
	enabled := true
	for _, s := range reg.FindAllStringSubmatch(input, -1) {
		if s[0] == "do()" {
			enabled = true
		} else if s[0] == "don't()" {
			enabled = false
		} else {
			if enabled {
				result += utils.ToInt(s[1]) * utils.ToInt(s[2])
			}
		}
	}
	return result
}

func main() {
	file, err := os.ReadFile("2024/day3/input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 3.1:", findUncorupptedWithDo(string(file), false))
	fmt.Println("Day 3.2:", findUncorupptedWithDo(string(file), true))
}
