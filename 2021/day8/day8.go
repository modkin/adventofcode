package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getArray(segments map[string]int, input string) [7]int {
	var output [7]int
	for _, r := range []rune(input) {
		output[segments[string(r)]] = 1
	}
	return output
}

func removeString(in string, remove string) string {
	for _, r := range []rune(remove) {
		in = strings.ReplaceAll(in, string(r), "")
	}
	return in
}

func main() {
	file, err := os.Open("2021/day8/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	input := make([][]string, 0)
	output := make([][]string, 0)
	display := make(map[[7]int]int, 0)
	display[[7]int{1, 1, 1, 0, 1, 1, 1}] = 0
	display[[7]int{0, 0, 1, 0, 0, 1, 0}] = 1
	display[[7]int{1, 0, 1, 1, 1, 0, 1}] = 2
	display[[7]int{1, 0, 1, 1, 0, 1, 1}] = 3
	display[[7]int{0, 1, 1, 1, 0, 1, 0}] = 4
	display[[7]int{1, 1, 0, 1, 0, 1, 1}] = 5
	display[[7]int{1, 1, 0, 1, 1, 1, 1}] = 6
	display[[7]int{1, 0, 1, 0, 0, 1, 0}] = 7
	display[[7]int{1, 1, 1, 1, 1, 1, 1}] = 8
	display[[7]int{1, 1, 1, 1, 0, 1, 1}] = 9

	for scanner.Scan() {
		inOut := strings.Split(scanner.Text(), "|")
		output = append(output, strings.Fields(inOut[1]))
		input = append(input, strings.Fields(inOut[0]))
	}
	total := 0
	for _, i := range output {
		for _, o := range i {
			if len(o) == 2 || len(o) == 3 || len(o) == 7 || len(o) == 4 {
				total++
			}
		}
	}
	fmt.Println(total)
	totalsum := 0
	for index, in := range input {
		var numbers [10]string
		letterCount := make(map[string]int)
		segments := make(map[string]string)
		for _, dig := range in {
			for _, d := range []rune(dig) {
				letterCount[string(d)]++
			}
			if len(dig) == 2 {
				numbers[1] = dig
			} else if len(dig) == 7 {
				numbers[8] = dig
			} else if len(dig) == 3 {
				numbers[7] = dig
			} else if len(dig) == 4 {
				numbers[4] = dig
			}
		}
		fmt.Println(letterCount)
		segments["t"] = removeString(numbers[7], numbers[1])
		for letter, count := range letterCount {
			if count == 4 {
				segments["bl"] = letter
			} else if count == 9 {
				segments["br"] = letter
			} else if count == 6 {
				segments["tl"] = letter
			} else if count == 8 {
				if letter != segments["t"] {
					segments["tr"] = letter
				}
			}
		}
		mid := strings.ReplaceAll(numbers[4], segments["tl"], "")
		mid = strings.ReplaceAll(mid, segments["tr"], "")
		mid = strings.ReplaceAll(mid, segments["br"], "")
		segments["m"] = mid
		for _, letter := range []rune(numbers[8]) {
			notFound := true
			for _, l := range segments {
				if l == string(letter) {
					notFound = false
				}
			}
			if notFound {
				segments["b"] = string(letter)
			}
		}
		fmt.Println(segments)
		// trun segments
		// map string to pos in display array
		segments2 := make(map[string]int)
		segments2[segments["t"]] = 0
		segments2[segments["tl"]] = 1
		segments2[segments["tr"]] = 2
		segments2[segments["m"]] = 3
		segments2[segments["bl"]] = 4
		segments2[segments["br"]] = 5
		segments2[segments["b"]] = 6

		var number string
		for _, s := range output[index] {
			number += strconv.Itoa(display[getArray(segments2, s)])
		}
		totalsum += utils.ToInt(number)
	}
	fmt.Println(totalsum)
}
