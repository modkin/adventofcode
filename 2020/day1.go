package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	numbers := make([]int, 0)
	file, err := os.Open("2020/input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next, _ := strconv.Atoi(scanner.Text())
		for _, elem := range numbers {
			if next+elem == 2020 {
				fmt.Println(next * elem)
			}
		}
		numbers = append(numbers, next)
	}
}
