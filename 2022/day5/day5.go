package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2022/day5/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	stackNbr := 9
	stacks := make([][]string, stackNbr)
	for i := 0; i < stackNbr; i++ {
		stacks[i] = make([]string, 0)
	}
	moves := make([][3]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || string(line[1]) == "1" {
			continue
		}
		if string(line[0]) == " " || string(line[0]) == "[" {
			for i := 0; i < len(line); i += 4 {
				//re := regexp.MustCompile(`\w*`)
				part := line[i : i+3]
				//char := re.FindString(line[i*4 : (i+1)*4])
				if " " != string(part[2]) {
					stacks[i/4] = append(stacks[i/4], string(part[1]))
				}
			}
		} else {
			split := strings.Split(line, " ")
			moves = append(moves, [3]int{utils.ToInt(split[1]), utils.ToInt(split[3]), utils.ToInt(split[5])})
		}
	}

	fmt.Println(stacks)
	fmt.Println(moves)

	for _, move := range moves {
		amount := move[0]
		from := move[1] - 1
		to := move[2] - 1
		for i := 0; i < amount; i++ {
			stacks[to] = append([]string{stacks[from][0]}, stacks[to]...)
			stacks[from] = stacks[from][1:]
		}
		fmt.Println(stacks)
		fmt.Println("Next")
	}

	for i := 0; i < 9; i++ {
		fmt.Print(stacks[i][0])
	}
	fmt.Println()

	fmt.Println("Day 3.1:", stacks)

}
