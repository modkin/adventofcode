package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(input string) int {
	current := 0
	for _, ru := range input {

		current += int(ru)
		current *= 17
		current = current % 256
	}
	return current
}

func main() {
	file, err := os.Open("2023/day15/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	//fmt.Println(hash("HASH"))
	sum := 0
	for _, str := range strings.Split(lines[0], ",") {
		sum += hash(str)
	}
	fmt.Println(sum)
}
