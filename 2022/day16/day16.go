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
	me      string
	elefant string
	timeMe  int
	timeEle int
}

//func findMaxState(input []state) state {
//	max := 0
//	var s state
//	for _, i2 := range input {
//		if i2.time <= 26 {
//			if i2.flow > max {
//				max = i2.flow
//				s = i2
//			}
//		}
//	}
//	return s
//}

//func isUseless(new state, states []state) bool {
//	for _, s := range states {
//		if new.me == s.me {
//			if new.flow < s.flow && new.time > s.time {
//				return true
//			}
//		}
//	}
//	return false
//}

//func isUseless2(new state, states map[string][]state) bool {
//	combinedPos := strings.Join([]string{new.me, new.elefant}, "-")
//	for _, s := range states[combinedPos] {
//		if new.flow < s.flow && new.time > s.time {
//			return true
//		}
//	}
//	combinedPos = strings.Join([]string{new.elefant, new.me}, "-")
//	for _, s := range states[combinedPos] {
//		if new.flow < s.flow && new.time > s.time {
//			return true
//		}
//	}
//	return false
//}

func getCheckSum(s state) string {
	ret := strings.Join(s.open, "-")
	ret += "-" + s.me + "-" + s.elefant + "-"
	ret += strconv.Itoa(s.timeMe) + strconv.Itoa(s.timeEle)
	ret += "-" + strconv.Itoa(s.flow)
	return ret
}

type search struct {
	path []string
	dist int
}

type target struct {
	tar  string
	dist int
}

func run2() {

	file, err := os.Open("/Users/dominikthoennes/go/src/adventofcode/2022/day16/input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	allValves := make(map[string]valve)
	//allState := make([]state, 0)
	//alreadyChecked := make(map[string]bool)
	//stateAtPos := make(map[string][]state)
	shortestPath := make(map[string]int)
	//distToValve := make(map[string]target)
	realValves := make([]string, 0)
	const maxTime = 30
	maxFlow := 0

	re := regexp.MustCompile(`\d*;`)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		name := line[1]
		flow := utils.ToInt(strings.Trim(re.FindString(scanner.Text()), ";"))
		if flow > 0 {
			realValves = append(realValves, name)
		}
		tunnel := make([]string, 0)
		for _, s := range line[9:] {
			tunnel = append(tunnel, strings.Trim(s, ","))
		}
		allValves[name] = valve{flow, tunnel}
	}
	allSearches := []search{search{[]string{"AA"}, 0}}
	for _, realVal := range realValves {
		allSearches = append(allSearches, search{[]string{realVal}, 0})
	}
	for len(allSearches) > 0 {
		newAllSearches := make([]search, 0)
		for _, curSearch := range allSearches {
			for _, t := range allValves[curSearch.path[len(curSearch.path)-1]].tunnel {
				if !utils.SliceContains(curSearch.path, t) {
					newSearch := search{append(utils.CopyStringSlice(curSearch.path), t), curSearch.dist + 1}
					newAllSearches = append(newAllSearches, newSearch)
					if utils.SliceContains(realValves, t) {
						conStr := []string{curSearch.path[0], t}
						sort.Strings(conStr)
						con := strings.Join(conStr, "-")
						con2 := strings.Join([]string{conStr[1], conStr[0]}, "-")
						if val, ok := shortestPath[con]; !ok || curSearch.dist+1 < val {
							shortestPath[con] = curSearch.dist + 1
							shortestPath[con2] = curSearch.dist + 1
						}
					}
				}
			}
		}
		allSearches = newAllSearches
	}
	//for con, dist := range shortestPath {
	//	split := strings.Split(con, "-")
	//	distToValve[split[0]] = target{split[1], dist}
	//	distToValve[split[1]] = target{split[0], dist}
	//}

	allStates := []state{state{[]string{"AA"}, 0, "AA", "AA", 1, 1}}
	for len(allStates) > 0 {
		newAllStates := make([]state, 0)
		for _, curState := range allStates {
			for _, mytargetValv := range realValves {
				if !utils.SliceContains(curState.open, mytargetValv) {
					dist := shortestPath[strings.Join([]string{curState.me, mytargetValv}, "-")]
					newTime := curState.timeMe + dist
					if newTime < maxTime {
						newFlow := curState.flow + (maxTime-newTime)*allValves[mytargetValv].flow
						if newFlow > maxFlow {
							maxFlow = newFlow
						}
						newState := state{
							open:    append(utils.CopyStringSlice(curState.open), mytargetValv),
							flow:    newFlow,
							me:      mytargetValv,
							elefant: curState.elefant,
							timeMe:  newTime + 1,
							timeEle: curState.timeEle,
						}
						//for _, eleTargetValv := range realValves {
						//	if !utils.SliceContains(curState.open, eleTargetValv) {
						//		dist = shortestPath[strings.Join([]string{newState.elefant, eleTargetValv}, "-")]
						//		newTime = curState.timeEle + dist
						//		if newTime < maxTime {
						//			newFlowEle := newState.flow + (maxTime-newTime)*allValves[eleTargetValv].flow
						//			if newFlowEle > maxFlow {
						//				maxFlow = newFlowEle
						//			}
						//			newState.open = append(newState.open, eleTargetValv)
						//			newState.flow = newFlowEle
						//			newState.elefant = eleTargetValv
						//			newState.timeEle = newTime
						//		}
						//	}
						//}
						newAllStates = append(newAllStates, newState)
					}
				}
			}

		}
		allStates = newAllStates
	}
	fmt.Println(maxFlow)

}

func main() {
	run2()
}
