package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2022/day5/testinput")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var stackNbr int
	for scanner.Scan() {
		if scanner.Text()[1] == '1' {
			line := strings.Split(scanner.Text(), " ")
			stackNbr = utils.ToInt(line[len(line)-1])
			break
		}
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)

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
				part := line[i : i+3]
				if " " != string(part[2]) {
					stacks[i/4] = append(stacks[i/4], string(part[1]))
				}
			}
		} else {
			split := strings.Split(line, " ")
			moves = append(moves, [3]int{utils.ToInt(split[1]), utils.ToInt(split[3]), utils.ToInt(split[5])})
		}
	}

	stackBackup := make([][]string, stackNbr)
	for i, stack := range stacks {
		stackBackup[i] = utils.CopyStringSlice(stack)
	}

	for _, move := range moves {
		amount := move[0]
		from := move[1] - 1
		to := move[2] - 1
		for i := 0; i < amount; i++ {
			stacks[to] = append([]string{stacks[from][0]}, stacks[to]...)
			stacks[from] = stacks[from][1:]
		}
	}

	fmt.Print("Day 5.1: ")
	for i := 0; i < stackNbr; i++ {
		fmt.Print(stacks[i][0])
	}
	fmt.Println()

	stacks = stackBackup

	for _, move := range moves {
		amount := move[0]
		from := move[1] - 1
		to := move[2] - 1
		stacks[to] = append(utils.CopyStringSlice(stacks[from][0:amount]), stacks[to]...)
		stacks[from] = stacks[from][amount:]
	}

	fmt.Print("Day 5.2: ")
	for i := 0; i < stackNbr; i++ {
		fmt.Print(stacks[i][0])
	}
	fmt.Println()

}
