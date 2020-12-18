package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func solveBlock(line string) string {

	for {
		closeBrace := strings.Index(line, ")")
		if closeBrace == -1 {
			break
		}
		openBrace := strings.LastIndex(line[:closeBrace], "(")
		line = line[:openBrace] + solveBlock(line[openBrace+1:closeBrace]) + line[closeBrace+1:]
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += utils.ToInt(solveBlock(line))
	}
	fmt.Println(sum)
}
