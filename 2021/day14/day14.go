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
	lettersCount := make(map[string]int)
	forumla := make(map[string]int)
	scanner.Scan()
	start = strings.Split(scanner.Text(), "")
	for i, s := range start {
		lettersCount[s]++
		if i < len(start)-1 {
			forumla[strings.Join([]string{start[i], start[i+1]}, "")] += 1

		}
	}
	scanner.Scan()

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		inserts[line[0]] = line[1]
		letters[line[1]] = true
	}

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
	fmt.Println("Day 14.1:", max-min)
	for step := 0; step < 40; step++ {
		newforumla := make(map[string]int)
		for key, count := range forumla {
			split := strings.Split(key, "")
			insert := inserts[key]
			lettersCount[insert] += count
			first := strings.Join([]string{split[0], insert}, "")
			second := strings.Join([]string{insert, split[1]}, "")
			newforumla[first] += count
			newforumla[second] += count
		}
		forumla = newforumla
	}
	max = 0
	min = math.MaxInt
	for _, count := range lettersCount {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	fmt.Println("Day 14.2:", max-min)
}
