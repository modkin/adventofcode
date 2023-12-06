package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2023/day6/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	races := make([][]int, 0)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	for i, time := range strings.Fields(lines[0])[1:] {
		distance := utils.ToInt(strings.Fields(lines[1])[i+1])
		races = append(races, []int{utils.ToInt(time), distance})
	}
	fmt.Println(races)

	var possiblilities []int
	for _, race := range races {
		counter := 0
		time := race[0]
		distance := race[1]
		for i := 1; i < distance; i++ {
			remainingTime := time - i
			if i*remainingTime > distance {
				counter++
			}
		}
		possiblilities = append(possiblilities, counter)
	}
	prod := 1
	for _, i2 := range possiblilities {
		prod *= i2
	}
	fmt.Println(prod)

}
