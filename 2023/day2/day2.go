package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2023/day2/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	var possibleIds []int
	var powers []int
	for _, line := range lines {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		items := strings.Split(line, ":")
		id := utils.ToInt(strings.Split(items[0], " ")[1])
		split := strings.Split(items[1], " ")
		GamePossible := true
		for i2, s := range split {

			if strings.Contains(s, "red") {
				number := utils.ToInt(split[i2-1])
				if number > maxRed {
					maxRed = number
				}
				if utils.ToInt(split[i2-1]) > 12 {
					GamePossible = false

				}
			}
			if strings.Contains(s, "green") {
				number := utils.ToInt(split[i2-1])
				if number > maxGreen {
					maxGreen = number
				}
				if utils.ToInt(split[i2-1]) > 13 {
					GamePossible = false

				}
			}
			if strings.Contains(s, "blue") {
				number := utils.ToInt(split[i2-1])
				if number > maxBlue {
					maxBlue = number
				}
				if utils.ToInt(split[i2-1]) > 14 {
					GamePossible = false

				}
			}

		}
		//if red <= 12 && green <= 13 && blue <= 14 {
		//	possibleIds = append(possibleIds, id)
		//}
		if GamePossible {
			possibleIds = append(possibleIds, id)
			fmt.Println(line)
		}
		GamePossible = true

		powers = append(powers, maxGreen*maxBlue*maxRed)
		maxBlue = 0
		maxRed = 0
		maxGreen = 0

	}
	//fmt.Println(possibleIds)
	sum := utils.SumSlice(possibleIds)
	fmt.Println("Day 1.1:", sum)
	fmt.Println(powers)
	sum2 := utils.SumSlice(powers)
	fmt.Println("Day 1.1:", sum2)

}
