package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day19/input"))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		fmt.Println(line)
	}
}
