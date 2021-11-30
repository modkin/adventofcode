package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2015/day15/testinput.txt")
	if err != nil {
		panic(err)
	}

	happinessMap := make(map[string][5]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		happinessMap[strings.Trim(lineSplit[0], ":")] = [5]int{
			utils.ToInt(strings.Trim(lineSplit[2], ",")),
			utils.ToInt(strings.Trim(lineSplit[4], ",")),
			utils.ToInt(strings.Trim(lineSplit[6], ",")),
			utils.ToInt(strings.Trim(lineSplit[8], ",")),
			utils.ToInt(strings.Trim(lineSplit[10], ","))}
	}
	fmt.Println(happinessMap)

}
