package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func calc(first int, second int, operator string) int {
	if operator == "*" {
		return first * second
	} else {
		return first + second
	}
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day18/testinput"))

	sum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		operator := false
		operators := make([]string, 0)
		braceCount := 0
		braces := make([]int, 20)

		popOperator := func() string {
			ret := operators[len(operators)-1]
			operators = operators[:len(operators)-1]
			return ret
		}
		first := true

		for _, elem := range line {
			if operator {
				operators = append(operators, elem)
				operator = false
			} else {
				number := strings.Split(elem, "")
				openB := strings.Count(elem, "(")
				closeB := strings.Count(elem, ")")
				if openB > 0 {
					for i := 0; i < openB; i++ {
						braceCount++
						braces[braceCount] = utils.ToInt(number[len(number)-1])
					}
				} else if closeB > 0 {
					for i := 0; i < closeB; i++ {
						if i == 0 {
							braces[braceCount] = calc(braces[braceCount], utils.ToInt(number[0]), popOperator())
						}
						braceCount--
						if len(operators) == 0 {
							braces[braceCount] = braces[braceCount+1]
						} else {
							braces[braceCount] = calc(braces[braceCount], braces[braceCount+1], popOperator())
						}
					}
				} else {
					if first {
						braces[braceCount] = utils.ToInt(number[0])
					} else {
						braces[braceCount] = calc(braces[braceCount], utils.ToInt(number[0]), popOperator())
					}

				}
				operator = true
			}
			first = false
		}
		sum += braces[0]
		fmt.Println(braces[0])
	}
	fmt.Println(sum)
}
