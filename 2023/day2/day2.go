package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("2023/day2/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	gameList := make([]map[string][]int, 0)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ":")
		GameSplit := strings.Split(strings.Trim(items[1], " "), " ")
		cubeMap := make(map[string][]int)
		for i2, s := range GameSplit {
			s = strings.Trim(s, ",;")
			if !utils.StringIsInt(s) {
				number := utils.ToInt(GameSplit[i2-1])
				cubeMap[s] = append(cubeMap[s], number)
			}
		}
		gameList = append(gameList, cubeMap)
	}

	possibleIdSum := 0
	powerSum := 0
	for i, cubeMap := range gameList {
		maxRed := slices.Max(cubeMap["red"])
		maxBlue := slices.Max(cubeMap["blue"])
		maxGreen := slices.Max(cubeMap["green"])
		if maxRed <= 12 && maxGreen <= 13 && maxBlue <= 14 {
			possibleIdSum += i + 1
		}
		powerSum += maxRed * maxGreen * maxBlue
	}

	fmt.Println("Day 2.1:", possibleIdSum)
	fmt.Println("Day 2.2:", powerSum)
}
