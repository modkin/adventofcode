package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func printList(input []string) {
	for _, s := range input {
		fmt.Print(s)
	}
	fmt.Println()
}

type pair struct {
	left      int
	right     int
	leftPair  *pair
	rightPair *pair
}

func explode(input []string) ([]string, bool) {
	nested := 0
	idx := 0
	exploded := false
	for i, r := range input {
		if r == "[" {
			nested++
		} else if r == "]" {
			nested--
		}
		if nested == 5 {
			idx = i
			exploded = true
			break
		}
	}
	if !exploded {
		return input, false
	}
	left, err := strconv.Atoi(input[idx+1])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(input[idx+3])
	if err != nil {
		panic(err)
	}
	for i := idx; i > 0; i-- {
		if number, err := strconv.Atoi(input[i]); err == nil {
			input[i] = strconv.Itoa(number + left)
			break
		}
	}
	for i := idx + 4; i < len(input); i++ {
		if number, err := strconv.Atoi(input[i]); err == nil {
			input[i] = strconv.Itoa(number + right)
			break
		}
	}
	output := append(utils.CopyStringSlice(input[:(idx)]), "0")
	output = append(output, input[(idx+5):]...)
	return output, true
}

func split(input []string) ([]string, bool) {
	idx := 0
	number := 0
	splited := false
	for i, r := range input {
		if n, err := strconv.Atoi(r); err == nil {
			if n > 9 {
				number = n
				idx = i
				splited = true
				break
			}
		}
	}
	if !splited {
		return input, false
	}
	left := number / 2
	right := (number + 1) / 2
	newPair := []string{"[", strconv.Itoa(left), ",", strconv.Itoa(right), "]"}
	output := append(utils.CopyStringSlice(input[:idx]), newPair...)
	output = append(output, input[(idx+1):]...)
	return output, true
}

func process(input []string) []string {
	changed := true
	for changed {
		input, changed = explode(input)
		if !changed {
			input, changed = split(input)
		}
	}
	return input
}

func magnitude(input []string) int {
	idx := 0
	for len(input) > 1 {
		tmp := 0
		for i, s := range input {
			if number, err := strconv.Atoi(s); err == nil {
				tmp = number * 3
				idx = i
				if number, err = strconv.Atoi(input[i+2]); err == nil {
					tmp += number * 2
					break
				}
			}
		}
		if len(input) == 5 {
			return tmp
		}
		tmpInput := append(utils.CopyStringSlice(input[:idx-1]), strconv.Itoa(tmp))
		tmpInput = append(tmpInput, input[idx+4:]...)
		input = utils.CopyStringSlice(tmpInput)
	}
	return math.MaxInt
}

func add(first []string, second []string) []string {
	line := append([]string{"["}, utils.CopyStringSlice(first)...)
	line = append(line, ",")
	line = append(line, utils.CopyStringSlice(second)...)
	line = append(line, "]")
	line = process(line)
	return line
}

func main() {
	file, err := os.Open("2021/day18/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	line := strings.Split(scanner.Text(), "")
	line = process(line)
	allLines := [][]string{line}
	for scanner.Scan() {
		newLine := strings.Split(scanner.Text(), "")
		allLines = append(allLines, utils.CopyStringSlice(newLine))
		line = add(line, newLine)

	}
	maxMag := 0
	for _, one := range allLines {
		for _, two := range allLines {
			if magnitude(one) == magnitude(two) {
				continue
			}
			tmpMag := magnitude(add(one, two))
			if tmpMag > maxMag {
				maxMag = tmpMag
			}
		}
	}
	//printList(line)
	fmt.Println("Day 18.1:", magnitude(line))
	fmt.Println("Day 18.2:", maxMag)
}
