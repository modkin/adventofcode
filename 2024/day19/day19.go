package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	towels := []string{}
	patterns := []string{}
	lines := utils.ReadFileIntoLines("2024/day19/input")

	towels = strings.Split(lines[0], ", ")
	patterns = lines[2:]

	var countCombs func(pat string) int
	var hashed = make(map[string]int)

	countCombs = func(pat string) int {
		if len(pat) == 0 {
			return 1
		}
		if _, ok := hashed[pat]; ok {
			return hashed[pat]
		}

		ret := 0
		for _, tow := range towels {
			if strings.HasPrefix(pat, tow) {
				ret += countCombs(pat[len(tow):])
			}
		}
		hashed[pat] = ret
		return ret
	}

	part2 := 0
	part1 := 0
	for _, pat := range patterns {
		hashed = make(map[string]int)
		combs := countCombs(pat)
		if combs > 0 {
			part1 += 1
		}
		part2 += combs
	}
	fmt.Println("Day 19.1:", part1)
	fmt.Println("Day 19.2:", part2)
}
