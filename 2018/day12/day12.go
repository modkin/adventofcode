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

func performEvolution(state []rune, updates map[[5]rune]rune) bool {
	tmp := make([]rune, len(state))
	copy(tmp, state)
	for i := 2; i < len(state)-2; i++ {
		update := [5]rune{state[i-2], state[i-1], state[i], state[i+1], state[i+2]}
		targetState := updates[update]
		tmp[i] = targetState
	}
	same := checkEqual(state, tmp)
	copy(state, tmp)
	return same
}

func checkEqual(one []rune, two []rune) bool {
	for i := 0; i < len(one)-2; i++ {
		if one[i] != two[i+1] {
			return false
		}
	}
	return true
}

func runXGenerations(count int) int {
	initstate, rules := parseInput()
	bufsize := int(len(initstate) / 2)
	buffer := make([]rune, bufsize)
	for idx, _ := range buffer {
		buffer[idx] = '.'
	}
	initstate = append(buffer, initstate...)
	initstate = append(initstate, buffer...)

	idx := 1
	for ; idx <= count; idx++ {
		same := performEvolution(initstate, rules)
		if same {
			break
		}
		if initstate[len(initstate)-3] == rune('#') {
			initstate = append(initstate, buffer...)
		}
	}
	sum := 0
	nonZero := 0
	for idx, _ := range initstate {
		if initstate[idx] == rune('#') {
			sum += idx - bufsize
			nonZero++
		}
	}
	if idx != count {
		sum += (count - idx) * nonZero
	}
	return sum
}

func Task1() {
	fmt.Println(runXGenerations(20))
}

func Task2() {
	fmt.Println(runXGenerations(50000000000))
}
