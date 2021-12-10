package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func checkMatch(left string, right string) bool {
	if left == "(" && right == ")" {
		return true
	}
	if left == "[" && right == "]" {
		return true
	}
	if left == "{" && right == "}" {
		return true
	}
	if left == "<" && right == ">" {
		return true
	}
	return false
}

func main() {

	file, err := os.Open("2021/day10/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	braceStack := make([]string, 0)
	illegal := make([]string, 0)
	incompleteLines := make([][]string, 0)
	for scanner.Scan() {
		braceStack = make([]string, 0)
		line := strings.Split(scanner.Text(), "")
		//fmt.Println(line)
		notIllegal := true
		for _, l := range line {
			if l == "(" || l == "[" || l == "{" || l == "<" {
				braceStack = append(braceStack, l)
			} else {
				if !checkMatch(braceStack[len(braceStack)-1], l) {
					//fmt.Println("expected", braceStack[len(braceStack)-1], "got", l)
					illegal = append(illegal, l)
					notIllegal = false
					break
				}
				braceStack = braceStack[:len(braceStack)-1]
			}
		}
		if notIllegal {
			incompleteLines = append(incompleteLines, braceStack)
		}
	}
	sum := 0
	for _, i := range illegal {
		if i == ")" {
			sum += 3
		} else if i == "]" {
			sum += 57
		} else if i == "}" {
			sum += 1197
		} else if i == ">" {
			sum += 25137
		}
	}
	fmt.Println("Day 10.1:", sum)
	scores := make([]int, 0)
	for _, line := range incompleteLines {
		score := 0
		for i := len(line) - 1; i >= 0; i-- {
			score *= 5
			switch line[i] {
			case "(":
				score += 1
			case "[":
				score += 2
			case "{":
				score += 3
			case "<":
				score += 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Println("Day 10.2", scores[len(scores)/2])
}
