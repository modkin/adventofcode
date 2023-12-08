package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmOfSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

func main() {
	file, err := os.Open("2023/day8/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	inst := make(map[string][]string)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	repeat := strings.Split(lines[0], "")
	for _, line := range lines[2:] {
		regex := regexp.MustCompile(`\w+`)
		split := regex.FindAllString(line, -1)
		inst[split[0]] = []string{split[1], split[2]}
	}

	sim := func(start string, end *regexp.Regexp) int {
		idx := 0
		counter := 0
		current := start
		for !end.MatchString(current) {
			counter++
			if repeat[idx] == "L" {
				current = inst[current][0]
			} else {
				current = inst[current][1]
			}
			idx++
			idx = idx % len(repeat)
		}
		return counter
	}

	fmt.Println("Day 8.1:", sim("AAA", regexp.MustCompile(`ZZZ`)))

	ghosts := make([]string, 0)
	for key := range inst {
		if key[2] == 'A' {
			ghosts = append(ghosts, key)
		}
	}

	arrive := make([]int, 0)
	for _, ghost := range ghosts {
		arrive = append(arrive, sim(ghost, regexp.MustCompile(`..Z`)))
	}
	fmt.Println("Day 8.2:", lcmOfSlice(arrive))

}
