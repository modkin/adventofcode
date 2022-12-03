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
	rucksacks2 := make([]string, 0)

	for scanner.Scan() {
		rucksacks2 = append(rucksacks2, scanner.Text())
		tmp := []rune(scanner.Text())
		firstHalf := tmp[0 : len(tmp)/2]
		secondHalf := tmp[len(tmp)/2:]
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

	groups := make([]rune, 0)
	for i := 0; i < len(rucksacks); i += 3 {
		var longest string
		if len(rucksacks2[i]) >= len(rucksacks2[i+1]) && len(rucksacks2[i]) >= len(rucksacks2[i+2]) {
			longest = rucksacks2[i]
		} else if len(rucksacks2[i+1]) >= len(rucksacks2[i]) && len(rucksacks2[i+1]) >= len(rucksacks2[i+2]) {
			longest = rucksacks2[i+1]
		} else {
			longest = rucksacks2[i+2]
		}
		fmt.Println("NEW")
		fmt.Println(rucksacks2[i])
		fmt.Println(rucksacks2[i+1])
		fmt.Println(rucksacks2[i+2])
		fmt.Println(longest)
		for _, m := range longest {
			if strings.Contains(rucksacks2[i], string(m)) && strings.Contains(rucksacks2[i+1], string(m)) && strings.Contains(rucksacks2[i+2], string(m)) {
				groups = append(groups, m)
				break
			}
		}
	}
	fmt.Println("Day 3.1:", sum)
	sum = 0
	for i := 0; i < len(groups); i++ {
		tmp := groups[i]
		if tmp >= 97 {
			sum += int(groups[i]) - 96
		} else {
			sum += int(groups[i]) - 38
		}
		//sum += doubles[i] - 64
	}

	fmt.Println("Day 3.2:", sum)

}
