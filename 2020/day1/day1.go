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
	var task1, task2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next, _ := strconv.Atoi(scanner.Text())
		for _, elem := range numbers {
			if next+elem == 2020 {
				task1 = next * elem
			}
			for _, elem2 := range numbers {
				if next+elem+elem2 == 2020 {
					task2 = next * elem * elem2
				}
			}
		}
		numbers = append(numbers, next)
	}
	fmt.Println("Task 1.1:", task1)
	fmt.Println("Task 1.2:", task2)
}
