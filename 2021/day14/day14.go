package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2021/day14/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	start := make([]string, 1)
	inserts := make(map[string]string)
	letters := make(map[string]bool)
	scanner.Scan()
	start = strings.Split(scanner.Text(), "")
	scanner.Scan()

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		inserts[line[0]] = line[1]
		letters[line[1]] = true
	}
	fmt.Println(start)
	fmt.Println(inserts)

	for step := 0; step < 10; step++ {
		startCopy := make([]string, 0)
		for i := 0; i < len(start)-1; i++ {
			current := strings.Join([]string{start[i], start[i+1]}, "")
			startCopy = append(startCopy, start[i])
			startCopy = append(startCopy, inserts[current])
			//start = startCopy
		}
		startCopy = append(startCopy, start[len(start)-1])

		start = utils.CopyStringSlice(startCopy)
		//fmt.Println(start)
	}
	max := 0
	min := math.MaxInt
	for l := range letters {
		count := utils.CountStringinStringSlice(start, l)
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	fmt.Println(max - min)
}
