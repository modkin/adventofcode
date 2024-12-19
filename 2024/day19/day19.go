package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func appendIfPossible(cur string, pattern string, towels []string) map[string]bool {
	out := make(map[string]bool)
	for _, tow := range towels {
		if len(cur+tow) <= len(pattern) && cur+tow == pattern[0:len(cur+tow)] {
			out[cur+tow] = true
		}
	}
	return out
}

func main() {

	towels := []string{}
	patterns := []string{}

	lines := utils.ReadFileIntoLines("2024/day19/input")
	for i, line := range lines {
		if i == 0 {
			split := strings.Split(line, ", ")
			for _, s := range split {
				towels = append(towels, s)
			}
		} else {
			if line == "" {
				continue
			}
			patterns = append(patterns, line)
		}
	}

	possible := 0
outer:
	for i, pattern := range patterns {
		fmt.Println(i)
		allCur := appendIfPossible("", pattern, towels)
		for len(allCur) > 0 {
			newAllCur := make(map[string]bool)
			for s, _ := range allCur {
				for s2, _ := range appendIfPossible(s, pattern, towels) {
					newAllCur[s2] = true
				}

			}
			allCur = newAllCur
			for s, _ := range allCur {
				if s == pattern {
					possible++
					continue outer
				}
			}
		}

	}
	fmt.Println(possible)
}
