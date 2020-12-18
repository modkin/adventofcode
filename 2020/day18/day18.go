package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day18/input"))
	for scanner.Scan() {
		foobar := strings.Split(scanner.Text(), ",")
		fmt.Println(foobar)
	}
}
