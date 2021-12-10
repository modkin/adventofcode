package main

import (
	"bufio"
	"fmt"
	"os"
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

	file, err := os.Open("2021/day10/testinput")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	braceStack := make([]string, 0)
	illegal := make([]string, 0)
	incompleteLines := make([][]string, 0)
	for scanner.Scan() {
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
			incompleteLines = append(incompleteLines, line)
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
	fmt.Println(sum)
	fmt.Println(incompleteLines)
}
