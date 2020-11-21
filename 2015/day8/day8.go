package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2015/day8/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var stringlen, reallen int
	for scanner.Scan() {
		line := scanner.Text()
		stringlen += len(line)
		lineTrim := strings.TrimPrefix(line, `"`)
		lineTrim = strings.TrimSuffix(lineTrim, `"`)
		lineUnquote, _ := strconv.Unquote(`"` + lineTrim + `"`)
		reallen += len(lineUnquote)
	}
	fmt.Println("Task 8.1:", stringlen-reallen)
}
