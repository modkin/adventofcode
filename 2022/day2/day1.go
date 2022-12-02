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
	fmt.Println(rounds)

	points := 0
	for i := 0; i < len(rounds); i++ {
		points += pointMap[rounds[i][1]]
		you := int(rounds[i][1])
		other := int(rounds[i][0])
		if you == 'X' && other == 'A' || you == 'Y' && other == 'B' || you == 'Z' && other == 'C' {
			points += 3
		} else {
			if you == 'Y' && other == 'A' {
				points += 6
			} else if you == 'Z' && other == 'B' {
				points += 6
			} else if you == 'X' && other == 'C' {
				points += 6
			}
		}
	}

	pointTwoMap := map[rune]int{'X': 0, 'Y': 3, 'Z': 6}
	pointChooseMap := map[rune]int{'A': 1, 'B': 2, 'C': 3}
	pointsTwo := 0
	for i := 0; i < len(rounds); i++ {
		pointsTwo += pointTwoMap[rounds[i][1]]
		you := rounds[i][1]
		other := rounds[i][0]
		if you == 'X' {
			if other == 'A' {
				pointsTwo += pointMap['Z']
			} else if other == 'B' {
				pointsTwo += pointMap['X']
			} else if other == 'C' {
				pointsTwo += pointMap['Y']
			}
		} else if you == 'Y' {
			pointsTwo += pointChooseMap[other]
		} else if you == 'Z' {
			if other == 'A' {
				pointsTwo += pointMap['Y']
			} else if other == 'B' {
				pointsTwo += pointMap['Z']
			} else if other == 'C' {
				pointsTwo += pointMap['X']
			}
		}
	}

	fmt.Println("Day 1.1:", points)
	fmt.Println("Day 1.1:", pointsTwo)

}
