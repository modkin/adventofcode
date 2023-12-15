package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(input string) int {
	current := 0
	for _, ru := range input {

		current += int(ru)
		current *= 17
		current = current % 256
	}
	return current
}

func main() {
	file, err := os.Open("2023/day15/input")
	if err != nil {
		panic(err)
	}

	boxes := make(map[int][]string)
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	//fmt.Println(hash("HASH"))
	//sum := 0
	for _, str := range strings.Split(lines[0], ",") {
		if strings.Contains(str, "-") {
			split := strings.Split(str, "-")
			label := split[0]
			box := hash(label)
			lenses := boxes[box]
			for i, lens := range lenses {

				if strings.Fields(lens)[0] == label {
					boxes[box] = append(lenses[:i], lenses[(i+1):]...)
					break
				}
			}
		} else if strings.Contains(str, "=") {
			split := strings.Split(str, "=")
			label := split[0]
			box := hash(label)
			lenses := boxes[box]
			notFound := true
			for i, lens := range lenses {

				if strings.Fields(lens)[0] == label {
					old := strings.Fields(lens)[1]
					boxes[box][i] = strings.ReplaceAll(boxes[box][i], old, split[1])
					notFound = false
					break
				}
			}
			if notFound {
				boxes[box] = append(boxes[box], split[0]+" "+split[1])
			}
		}
	}
	focusSum := 0
	for boxNbr, lenses := range boxes {
		for slot, lens := range lenses {
			focus := (boxNbr + 1) * (slot + 1) * utils.ToInt(strings.Fields(lens)[1])
			focusSum += focus
		}

	}
	fmt.Println("Day 15.2: ", focusSum)
}
