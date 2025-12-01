package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFileIntoLines("2025/day1/input")
	fmt.Println(len(lines))

	dialPos := 50
	password := 0

	for _, line := range lines {
		amount := utils.ToInt(line[1:])
		if line[0] == 'L' {
			dialPos = dialPos - amount
		} else {
			dialPos = dialPos + amount
		}

		dialPos = dialPos % 100

		if dialPos < 0 {

			dialPos = 100 + dialPos
		}
		fmt.Println(line, dialPos)
		if dialPos == 0 {
			password++
		}
	}
	fmt.Println(password)
}
