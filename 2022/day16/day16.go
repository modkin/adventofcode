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

func getCheckSum(in state) string {
	pos := []string{in.me, in.elefant}
	sort.Strings(pos)
	ret := strings.Join(pos, "-")
	ret += strconv.Itoa(in.timeMe) + strconv.Itoa(in.timeEle)
	ret += strconv.Itoa(in.flow)
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
	shortestPath := make(map[string]int)
	alreadyChecked := make(map[string]bool)
	realValves := make([]string, 0)
	const maxTime = 26
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
						tmpState := state{
							open:    append(utils.CopyStringSlice(curState.open), mytargetValv),
							flow:    newFlow,
							me:      mytargetValv,
							elefant: curState.elefant,
							timeMe:  newTime + 1,
							timeEle: curState.timeEle,
						}
						for _, eleTargetValv := range realValves {
							if !utils.SliceContains(tmpState.open, eleTargetValv) {
								newState := state{
									open:    utils.CopyStringSlice(tmpState.open),
									flow:    tmpState.flow,
									me:      tmpState.me,
									elefant: tmpState.elefant,
									timeMe:  tmpState.timeMe,
									timeEle: tmpState.timeEle,
								}
								dist = shortestPath[strings.Join([]string{newState.elefant, eleTargetValv}, "-")]
								newTime = curState.timeEle + dist
								if newTime < maxTime {
									newFlowEle := newState.flow + (maxTime-newTime)*allValves[eleTargetValv].flow
									if newFlowEle > maxFlow {
										maxFlow = newFlowEle
										fmt.Println(maxFlow)
									}
									newState.open = append(newState.open, eleTargetValv)
									newState.flow = newFlowEle
									newState.elefant = eleTargetValv
									newState.timeEle = newTime + 1
								}
								csum := getCheckSum(newState)
								if _, ok := alreadyChecked[csum]; !ok {
									newAllStates = append(newAllStates, newState)
									alreadyChecked[csum] = true
								}
							}

						}

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
