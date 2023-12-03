package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("2023/day3/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	var allNumbers []int
	number := ""
	adjacent := false
	for lineIdx, line := range lines {
		if adjacent && number != "" {
			allNumbers = append(allNumbers, utils.ToInt(number))
		}
		number = ""
		adjacent = false
		for charIdx, char := range line {
			if unicode.IsDigit(char) {
				number += string(char)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x+lineIdx < 0 || x+lineIdx >= len(lines)-1 || y+charIdx < 0 || y+charIdx > len(lines[0])-1 {
							continue
						}
						if lines[lineIdx+x][charIdx+y] != '.' && !unicode.IsDigit(rune(lines[lineIdx+x][charIdx+y])) {
							adjacent = true
						}

					}
				}
			} else {
				if adjacent && number != "" {
					allNumbers = append(allNumbers, utils.ToInt(number))
				}
				number = ""
				adjacent = false
			}
		}
	}

	//fmt.Println(allNumbers)
	sum := utils.SumSlice(allNumbers)
	fmt.Println("Day 3.1:", sum)

	var gearRatios []int
	for lineIdx, line := range lines {
		for charIdx, char := range line {
			if char == '*' {
				var adjacentNumbers []int
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x+lineIdx < 0 || x+lineIdx >= len(lines) || y+charIdx < 0 || y+charIdx >= len(lines[0]) {
							continue
						}
						if unicode.IsDigit(rune(lines[lineIdx+x][charIdx+y])) {
							ytmp := y
							number = ""
							for i := -2; i <= 2; i++ {
								if i+charIdx < 0 || y+charIdx > len(lines[0])-1 {
									continue
								}

								adjacentChar := lines[lineIdx+x][charIdx+ytmp+i]
								if unicode.IsDigit(rune(adjacentChar)) {
									number += string(adjacentChar)
									if i >= 0 {
										y++
									}
								} else {
									if i < 0 && number != "" {
										number = ""
									}
									if i > 0 && number != "" {
										break
									}
								}
							}
							adjacentNumbers = append(adjacentNumbers, utils.ToInt(number))
						}

					}
				}
				fmt.Println(adjacentNumbers)
				if len(adjacentNumbers) == 2 {
					gearRatios = append(gearRatios, adjacentNumbers[0]*adjacentNumbers[1])
				}
			} else {
				if adjacent && number != "" {
					allNumbers = append(allNumbers, utils.ToInt(number))
				}
				number = ""
				adjacent = false
			}
		}
	}
	//fmt.Println(gearRatios)
	sum2 := utils.SumSlice(gearRatios)
	fmt.Println("Day 3.2:", sum2)
}
