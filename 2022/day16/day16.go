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
	me      string
	elefant string
	time    int
}

func findMaxState(input []state) state {
	max := 0
	var s state
	for _, i2 := range input {
		if i2.time <= 26 {
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
		if new.me == s.me {
			if new.flow < s.flow && new.time > s.time {
				return true
			}
		}
	}
	return false
}

func isUseless2(new state, states map[string][]state) bool {
	combinedPos := strings.Join([]string{new.me, new.elefant}, "-")
	for _, s := range states[combinedPos] {
		if new.flow < s.flow && new.time > s.time {
			return true
		}

	}
	return false
}

func getCheckSum(s state) string {
	ret := strings.Join(s.open, "-")
	ret += "-" + s.me + "-" + s.elefant + "-"
	ret += strconv.Itoa(s.time)
	ret += "-" + strconv.Itoa(s.flow)
	return ret
}

func run() {

	file, err := os.Open("/Users/dominikthoennes/go/src/adventofcode/2022/day16/testinput")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	allValves := make(map[string]valve)
	allState := make([]state, 0)
	alreadyChecked := make(map[string]bool)
	stateAtPos := make(map[string][]state)
	const maxTime = 26

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
	allState = append(allState, state{[]string{"AA"}, 0, "AA", "AA", 1})
	newOpen := true
	currentMax := 0
	for newOpen {
		newOpen = false
		newAllState := make([]state, 0)
		for _, curState := range allState {

			if curState.time >= maxTime {
				continue
			}
			newOpen = true

			for _, newOwnPos := range allValves[curState.me].tunnel {
				newState := state{utils.CopyStringSlice(curState.open), curState.flow, newOwnPos, curState.elefant, curState.time + 1}
				for _, newElefantPos := range allValves[curState.elefant].tunnel {
					newState.elefant = newElefantPos

					if !utils.SliceContains(newState.open, newOwnPos) && allValves[newOwnPos].flow != 0 {
						for _, newNewElefantPos := range allValves[newElefantPos].tunnel {
							newState2 := state{append(utils.CopyStringSlice(curState.open), newOwnPos), curState.flow, newOwnPos, newNewElefantPos, curState.time + 1}
							newState2.flow += (maxTime - newState2.time) * allValves[newOwnPos].flow
							newState2.time++
							newAllState = append(newAllState, newState2)
						}
					}
					if !utils.SliceContains(newState.open, newElefantPos) && allValves[newElefantPos].flow != 0 {
						for _, newNewOwnPos := range allValves[newOwnPos].tunnel {
							newState2 := state{append(utils.CopyStringSlice(curState.open), newElefantPos), curState.flow, newNewOwnPos, newElefantPos, curState.time + 1}
							newState2.flow += (maxTime - newState2.time) * allValves[newElefantPos].flow
							newState2.time++
							newAllState = append(newAllState, newState2)
						}
					}

					newAllState = append(newAllState, newState)
				}
			}

		}
		allState = make([]state, 0)
		for _, s := range newAllState {
			combinedPos := strings.Join([]string{s.me, s.elefant}, "-")
			if _, ok := stateAtPos[combinedPos]; ok {
				stateAtPos[combinedPos] = append(stateAtPos[s.me], s)
			} else {
				stateAtPos[combinedPos] = []state{s}
			}

		}
		for _, s := range newAllState {
			//if !isUseless(s, newAllState) {
			if !isUseless2(s, stateAtPos) {
				if _, ok := alreadyChecked[getCheckSum(s)]; !ok {
					allState = append(allState, s)
					alreadyChecked[getCheckSum(s)] = true
					if s.flow > currentMax {
						currentMax = s.flow
						fmt.Println(currentMax, len(allState))
					}
				}
			}
		}

	}
	fmt.Println(currentMax)
}

func main() {
	run()
}
