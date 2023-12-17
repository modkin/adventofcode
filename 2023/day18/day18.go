package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("2023/day18/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
}
