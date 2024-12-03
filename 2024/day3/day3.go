package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func findUncoruppted(input string) int {
	reg := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	found := reg.FindAllStringSubmatch(input, -1)
	result := 0
	for _, s := range found {
		fmt.Println(s[0], s[1], s[2])
		result += utils.ToInt(s[1]) * utils.ToInt(s[2])
	}
	return result
}

func main() {
	file, err := os.Open("2024/day3/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		result += findUncoruppted(scanner.Text())
	}

	fmt.Println(result)
}
