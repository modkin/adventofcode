package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("2021/day1/testinput")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(utils.ToInt(scanner.Text()))
	}
}
