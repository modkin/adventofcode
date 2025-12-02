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
	password2 := 0

	for _, line := range lines {
		amount := utils.ToInt(line[1:])
		dialOld := dialPos
		if line[0] == 'L' {
			dialPos = dialPos - amount
		} else {
			dialPos = dialPos + amount
		}
		if dialOld*dialPos < 0 {
			password2++
		}

		if utils.IntAbs(dialPos) > 100 {
			password2 += utils.IntAbs(dialPos / 100)
		}

		fmt.Println(line, dialPos, password2, password)
		dialPos = dialPos % 100

		//if dialPos < 0 {
		//	dialPos = 100 + dialPos
		//	dialPos = dialPos % 100
		//}
		if dialPos == 0 {
			password++
		}
	}
	fmt.Println(password)
	fmt.Println(password2 + password)

	dialPos = 50
	passBrute := 0
	for _, line := range lines {
		amount := utils.ToInt(line[1:])

		for i := 0; i < amount; i++ {
			if line[0] == 'L' {
				dialPos--
			} else {
				dialPos++
			}
			dialPos = dialPos % 100
			if dialPos == 0 {
				passBrute++
			}
		}
	}
	fmt.Println(passBrute)
}
