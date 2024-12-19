package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func appendIfPossible(cur string, pattern string, towels []string) map[string]int {
	out := make(map[string]int)
	for _, tow := range towels {
		join := strings.Join(strings.Split(cur, ","), "") + tow
		if len(join) <= len(pattern) && join == pattern[0:len(join)] {
			if cur == "" {
				out[tow] = 1
			} else {
				out[cur+","+tow] = 1
			}
		}
	}
	return out
}

func main() {

	towels := []string{}
	patterns := []string{}
	maxTowelLenght := 0

	lines := utils.ReadFileIntoLines("2024/day19/input")
	for i, line := range lines {
		if i == 0 {
			split := strings.Split(line, ", ")
			for _, s := range split {
				towels = append(towels, s)
				if maxTowelLenght < len(s) {
					maxTowelLenght = len(s)
				}
			}
		} else {
			if line == "" {
				continue
			}
			patterns = append(patterns, line)
		}
	}

	//possible := 0
	allPossible := make(map[string]int)
	//outer:
	for iOut, pattern := range patterns {
		fmt.Println(iOut)
		fmt.Println(pattern)
		allCur := appendIfPossible("", pattern, towels)
		for len(allCur) > 0 {
			newAllCur := make(map[string]int)
			for s, _ := range allCur {
				for s2, num := range appendIfPossible(s, pattern, towels) {
					newAllCur[s2] = num
				}

			}
			allCur = newAllCur
			for iPat := 0; iPat < len(pattern)-2; iPat++ {
				newAllCur = make(map[string]int)
				minPattern := pattern[0:iPat]
				for s, b := range allCur {
					newStr := ""
					split := strings.Split(s, ",")

					for i, s2 := range split {
						if len(newStr+s2) <= len(minPattern) && newStr+s2 == minPattern[0:len(newStr+s2)] {
							newStr += s2
							if i == len(split)-1 {
								newAllCur[newStr] += b
							}
						} else {
							if newStr != "" {
								newAllCur[newStr+","+strings.Join(split[i:], ",")] += b
								break
							} else {
								newAllCur[s] += b
								break
							}
						}
					}
				}
				allCur = newAllCur
			}

			for s, b := range allCur {
				join := strings.Join(strings.Split(s, ","), "")
				if join == pattern {
					allPossible[pattern] += b
				}
			}

		}

	}
	fmt.Println(allPossible)
	sum := 0
	for _, i := range allPossible {
		sum += i
	}
	fmt.Println(sum)
}
