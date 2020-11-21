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
	var stringlen, reallen, escapelen int
	for scanner.Scan() {
		line := scanner.Text()
		escapeLine := strings.ReplaceAll(line, `\`, `\\`)
		escapeLine = strings.ReplaceAll(escapeLine, `"`, `\"`)
		escapelen += len(escapeLine) + 2
		stringlen += len(line)
		lineTrim := strings.TrimPrefix(line, `"`)
		lineTrim = strings.TrimSuffix(lineTrim, `"`)
		lineUnquote, _ := strconv.Unquote(`"` + lineTrim + `"`)
		reallen += len(lineUnquote)
	}
	fmt.Println("Task 8.1:", stringlen-reallen)
	fmt.Println("Task 8.2:", escapelen-stringlen)
}
