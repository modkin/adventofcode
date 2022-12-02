package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2022/day2/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rounds := make([][2]rune, 0)
	pointMap := map[rune]int{'X': 1, 'Y': 2, 'Z': 3}

	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		rounds = append(rounds, [2]rune{[]rune(tmp[0])[0], []rune(tmp[1])[0]})
	}

	points := 0
	for i := 0; i < len(rounds); i++ {
		points += pointMap[rounds[i][1]]
		you := rounds[i][1]
		other := rounds[i][0]
		// distance is constant for draw
		if you-other == 23 {
			points += 3
		} else {
			if you == 'X' {
				you += 3
			}
			result := you - other - 23
			if result == 1 {
				points += 6
			}
		}
	}

	pointTwoMap := map[rune]int{'X': 0, 'Y': 3, 'Z': 6}
	pointsTwo := 0
	for i := 0; i < len(rounds); i++ {
		pointsTwo += pointTwoMap[rounds[i][1]]
		you := rounds[i][1]
		other := rounds[i][0]
		if you == 'X' {
			tmpPoints := pointMap[other+23] - 1
			if tmpPoints == 0 {
				tmpPoints = 3
			}
			pointsTwo += tmpPoints
		} else if you == 'Y' {
			pointsTwo += pointMap[other+23]
		} else if you == 'Z' {
			pointsTmp := pointMap[other+23] + 1
			if pointsTmp == 4 {
				pointsTmp = 1
			}
			pointsTwo += pointsTmp
		}
	}
	if points != 11603 || pointsTwo != 12725 {
		panic(err)
	}
	fmt.Println("Day 1.1:", points)
	fmt.Println("Day 1.1:", pointsTwo)

}
