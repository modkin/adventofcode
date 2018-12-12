package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() ([]rune, map[[5]rune]rune) {
	file, err := os.Open("day12/day12-input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstline := scanner.Text()
	initialString := strings.Split(firstline, " ")[2]
	initState := []rune(initialString)
	//skip empty line
	scanner.Scan()

	rules := make(map[[5]rune]rune)
	for scanner.Scan() {
		rule := scanner.Text()
		src := strings.Split(rule, " => ")[0]
		srcRune := []rune(src)
		srcArray := [5]rune{srcRune[0], srcRune[1], srcRune[2], srcRune[3], srcRune[4]}
		dst := strings.Split(rule, " => ")[1]
		dstRune := []rune(dst)
		rules[srcArray] = dstRune[0]
	}
	return initState, rules
}

func performEvolution(state []rune, updates map[[5]rune]rune) {
	tmp := make([]rune, len(state))
	copy(tmp, state)
	for i := 2; i < len(state)-2; i++ {
		update := [5]rune{state[i-2], state[i-1], state[i], state[i+1], state[i+2]}
		targetState := updates[update]
		tmp[i] = targetState
	}
	copy(state, tmp)
}

func runXGenerations(count int) int {
	initstate, rules := parseInput()
	bufsize := 4 + (count-1)*2
	buffer := make([]rune, bufsize)
	for idx, _ := range buffer {
		buffer[idx] = '.'
	}
	initstate = append(buffer, initstate...)
	initstate = append(initstate, buffer...)

	for i := 1; i <= count; i++ {
		performEvolution(initstate, rules)
		//fmt.Println(string(initstate))
	}
	sum := 0
	for idx, _ := range initstate {
		if initstate[idx] == rune('#') {
			sum += idx - bufsize
		}
	}
	return sum
}

func Task1() {
	fmt.Println(runXGenerations(20))
}

func Task2() {
	//fmt.Println(runXGenerations(50000000000))
}
