package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2022/day3/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rucksacks := make([][2]string, 0)

	for scanner.Scan() {
		tmp := []rune(scanner.Text())
		firstHalf := tmp[0 : len(tmp)/2]
		secondHalf := tmp[len(tmp)/2 : len(tmp)]
		rucksacks = append(rucksacks, [2]string{string(firstHalf), string(secondHalf)})
	}

	doubles := make([]rune, 0)
	for i := 0; i < len(rucksacks); i++ {
		for m := 0; m < len(rucksacks[i][0]); m++ {
			if strings.Contains(rucksacks[i][1], string(rucksacks[i][0][m])) {
				doubles = append(doubles, rune(rucksacks[i][0][m]))
				break
			}
		}
	}

	fmt.Println("Day 3.1:", string(doubles))
	sum := 0
	for i := 0; i < len(doubles); i++ {
		tmp := doubles[i]
		if tmp >= 97 {
			sum += int(doubles[i]) - 96
		} else {
			sum += int(doubles[i]) - 38
		}
		fmt.Println(int(doubles[i]))
		//sum += doubles[i] - 64
	}

	fmt.Println("Day 3.1:", sum)

}
