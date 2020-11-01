package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("2015/day1/input")
	if err != nil {
		panic(err)
	}

	floor := 0
	position := 1
	basementPos := 0
	for _, elem := range content {
		if elem == '(' {
			floor += 1
		} else if elem == ')' {
			floor += -1
		}
		if floor == -1 && basementPos == 0 {
			basementPos = position
		}
		position += 1
	}
	fmt.Println("Task 1.1:", floor)
	fmt.Println("Task 1.2:", basementPos)
}
