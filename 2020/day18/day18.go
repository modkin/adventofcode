package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AddFirst(line string) string {
	ret := make([]string, 0)
	for _, multiplication := range strings.Split(line, "*") {
		splitMult := strings.Split(strings.TrimSpace(multiplication), " ")
		current := utils.ToInt(splitMult[0])
		for i := 0; i < len(splitMult)-2; i += 2 {
			current += utils.ToInt(splitMult[i+2])
		}
		ret = append(ret, strconv.Itoa(current))
	}
	return strings.Join(ret, " * ")
}

func solveBlock(line string, partTwo bool) string {

	for {
		closeBrace := strings.Index(line, ")")
		if closeBrace == -1 {
			break
		}
		openBrace := strings.LastIndex(line[:closeBrace], "(")
		line = line[:openBrace] + solveBlock(line[openBrace+1:closeBrace], partTwo) + line[closeBrace+1:]
	}
	if partTwo {
		line = AddFirst(line)
	}

	splitLine := strings.Split(line, " ")
	current := utils.ToInt(splitLine[0])
	for i := 0; i < len(splitLine)-2; i += 2 {
		if splitLine[i+1] == "*" {
			current *= utils.ToInt(splitLine[i+2])
		} else {
			current += utils.ToInt(splitLine[i+2])
		}
	}
	return strconv.Itoa(current)
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day18/input"))

	sum, sum2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += utils.ToInt(solveBlock(line, false))
		sum2 += utils.ToInt(solveBlock(line, true))
	}
	fmt.Println("Task 18.1:", sum)
	fmt.Println("Task 18.2:", sum2)
}
