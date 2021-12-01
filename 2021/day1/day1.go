package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("2021/day1/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	height := 0
	//next := 0
	count := 0
	//meassures := make([]int,0)
	for scanner.Scan() {
		next := utils.ToInt(scanner.Text())
		if next > height {
			count++
		}
		height = next
	}
	fmt.Println(count - 1)
}
