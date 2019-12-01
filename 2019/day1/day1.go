package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var fuelRequirement1 = 0
	var fuelRequirement2 = 0

	file, err := os.Open("2019/day1/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		fuelRequirement1 += i/3 - 2
		for i > 0 {
			i = i/3 - 2
			if i < 0 {
				i = 0
			}
			fuelRequirement2 += i
		}

	}
	fmt.Println("Task 1.1: fuel requirement: ", fuelRequirement1)
	fmt.Println("Task 1.2: fuel requirement: ", fuelRequirement2)

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
