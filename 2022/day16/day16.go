package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	flow   int
	tunnel []string
}

type state struct {
	open    []string
	flow    int
	current string
	time    int
}

func findMaxState(input []state) state {
	max := 0
	var s state
	for _, i2 := range input {
		if i2.time <= 30 {
			if i2.flow > max {
				max = i2.flow
				s = i2
			}
		}
	}
	return s
}

func isUseless(new state, states []state) bool {
	for _, s := range states {
		if new.current == s.current {
			sameOpen := true
			for _, v := range new.open {
				sameOpen = sameOpen && utils.SliceContains(s.open, v)
			}
			if sameOpen {
				if new.flow < s.flow && new.time >= s.flow {
					return true
				}
			}
		}
	}
	return false
}

func isSame(new state, states []state) bool {
	for _, s := range states {
		if new.current == s.current {
			sameOpen := true
			for _, v := range new.open {
				sameOpen = sameOpen && utils.SliceContains(s.open, v)
			}
			if sameOpen {
				if new.flow == s.flow && new.time == s.time {
					return true
				}
			}
		}
	}
	return false
}

func letIfFlow(s state, allValves map[string]valve) int {
	if s.time > 30 {
		return 0
	}
	add := 0
	for _, s2 := range s.open {
		add += allValves[s2].flow
	}
	return add
}

func getCheckSum(s state) string {
	ret := strings.Join(s.open, "-")
	ret += "-" + s.current + "-"
	ret += strconv.Itoa(s.time)
	ret += "-" + strconv.Itoa(s.flow)
	return ret
}

func main() {

	file, err := os.Open("2022/day16/testinput")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	allValves := make(map[string]valve)
	allState := make([]state, 0)
	alreadyChecked := make(map[string]bool)

	re := regexp.MustCompile(`\d*;`)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		name := line[1]
		flow := re.FindString(scanner.Text())
		tunnel := make([]string, 0)
		for _, s := range line[9:] {
			tunnel = append(tunnel, strings.Trim(s, ","))
		}
		allValves[name] = valve{utils.ToInt(strings.Trim(flow, ";")), tunnel}

	}
	fmt.Println(allValves)
	allState = append(allState, state{[]string{"AA"}, 0, "AA", 1})
	newOpen := true
	currentMax := 0
	for newOpen {
		newOpen = false
		newAllState := make([]state, 0)
		for i, maxState := range allState {
			if maxState.time > 30 {
				continue
			}
			flow := letIfFlow(maxState, allValves)
			if maxState.flow+(30-maxState.time)*flow < currentMax {
				continue
			}

			add := 0
			for _, s2 := range maxState.open {
				add += allValves[s2].flow
			}
			allState[i].flow += add
			newOpen = true
			//}

			//maxIdx := findMaxState(allState)
			//maxState := allState[maxIdx]
			//maxState.time++

			//for _, maxState := range allState {

			for _, v := range allValves[maxState.current].tunnel {
				newState := state{utils.CopyStringSlice(maxState.open), maxState.flow, v, maxState.time + 1}
				if !utils.SliceContains(newState.open, v) {
					newState2 := state{utils.CopyStringSlice(maxState.open), maxState.flow, v, maxState.time + 1}

					newState2.open = append(newState2.open, v)
					newState2.time++
					newState2.flow += letIfFlow(newState2, allValves)
					newAllState = append(newAllState, newState2)
					//newState.flow = letIfFlow(newState, allValves)
					//newState.open = append(newState.open, v)
					//newState.time++

				}
				newAllState = append(newAllState, newState)
			}
		}
		allState = make([]state, 0)
		for _, s := range newAllState {
			if !isUseless(s, newAllState) {
				if !isSame(s, allState) {
					if _, ok := alreadyChecked[getCheckSum(s)]; !ok {
						allState = append(allState, s)
						alreadyChecked[getCheckSum(s)] = true
					}
				}
			}
		}
		max := findMaxState(allState).flow
		if max > currentMax {
			currentMax = max
		}
		fmt.Println(currentMax)
	}

	maxState := 0
	for _, s := range allState {
		if s.flow > maxState {
			maxState = s.flow
			fmt.Println(s.open)
		}
	}
	fmt.Println(maxState)
}
