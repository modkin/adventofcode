package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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
	path    []string
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
			if new.flow < s.flow && new.time > s.time {
				return true
			}
		}
	}
	return false
}

func isUseless2(new state, states map[string][]state) bool {
	for _, s := range states[new.current] {
		if new.flow < s.flow && new.time > s.time {
			return true
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

func run() {

	file, err := os.Open("/Users/dominikthoennes/go/src/adventofcode/2022/day16/input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	allValves := make(map[string]valve)
	allState := make([]state, 0)
	alreadyChecked := make(map[string]bool)
	stateAtPos := make(map[string][]state)

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
	allState = append(allState, state{[]string{"AA"}, 0, "AA", 1, make([]string, 0)})
	newOpen := true
	currentMax := 0
	for newOpen {
		newOpen = false
		newAllState := make([]state, 0)
		for _, maxState := range allState {
			max := findMaxState(newAllState).flow
			if max > currentMax {
				currentMax = max
				fmt.Println(currentMax, len(allState))
			}
			if maxState.time >= 30 {
				continue
			}
			newOpen = true
			//}

			//maxIdx := findMaxState(allState)
			//maxState := allState[maxIdx]
			//maxState.time++

			//for _, maxState := range allState {

			for _, v := range allValves[maxState.current].tunnel {
				newPath := append(utils.CopyStringSlice(maxState.path), v)
				newState := state{utils.CopyStringSlice(maxState.open), maxState.flow, v, maxState.time + 1, newPath}
				if !utils.SliceContains(newState.open, v) && allValves[v].flow != 0 {
					newState2 := state{utils.CopyStringSlice(maxState.open), maxState.flow, v, maxState.time + 1, newPath}
					tmp := append(newState2.open, v)
					sort.Strings(tmp)
					newState2.open = tmp
					newState2.flow += (30 - newState2.time) * allValves[v].flow
					newState2.time++
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
			if _, ok := stateAtPos[s.current]; ok {
				stateAtPos[s.current] = append(stateAtPos[s.current], s)
			} else {
				stateAtPos[s.current] = []state{s}
			}

		}
		for _, s := range newAllState {
			//if !isUseless(s, newAllState) {
			if !isUseless2(s, stateAtPos) {
				if _, ok := alreadyChecked[getCheckSum(s)]; !ok {
					allState = append(allState, s)
					alreadyChecked[getCheckSum(s)] = true
				}
			}
		}

	}
	fmt.Println(currentMax)
}

func main() {
	run()
}
