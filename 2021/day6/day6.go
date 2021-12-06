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
	scanner := bufio.NewScanner(file)
	//lines := make([]ventline, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, l := range line {
			fishes = append(fishes, utils.ToInt(l))
		}
	}
	for i := 0; i < 256; i++ {
		for i2, fish := range fishes {
			if fish == 0 {
				fishes = append(fishes, 8)
				fishes[i2] = 6
			} else {
				fishes[i2]--
			}
		}
		fmt.Println(i)
	}
	fmt.Println(len(fishes))
}
