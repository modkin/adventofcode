package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2021/day6/input")
	if err != nil {
		panic(err)
	}

	fishes := make([]int, 0)
	var birthday [9]int
	scanner := bufio.NewScanner(file)
	//lines := make([]ventline, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, l := range line {
			fishes = append(fishes, utils.ToInt(l))
			birthday[utils.ToInt(l)]++
		}
	}
	for i := 0; i < 80; i++ {
		for i2, fish := range fishes {
			if fish == 0 {
				fishes = append(fishes, 8)
				fishes[i2] = 6
			} else {
				fishes[i2]--
			}
		}
	}
	fmt.Println("Day 6.1:", len(fishes))
	for i := 0; i < 256; i++ {
		tmp := birthday[0]
		for i := 0; i <= 7; i++ {
			birthday[i] = birthday[i+1]
		}
		birthday[8] = tmp
		birthday[6] += tmp
	}
	sum := 0
	for _, b := range birthday {
		sum += b
	}
	fmt.Println("Day 6.2:", sum)
}
