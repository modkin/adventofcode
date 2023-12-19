package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type node struct {
	curWF  string
	ranges map[string][2]int
}

func copyString2IntMap(in map[string][2]int) map[string][2]int {
	ret := make(map[string][2]int)
	for s, i := range in {
		ret[s] = i
	}
	return ret
}

func main() {
	file, err := os.Open("2023/day19/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	workflows := make(map[string][]string)
	var parts []map[string]int
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	onRules := true
	for _, line := range lines {
		if line == "" {
			onRules = false
			continue
		}
		if onRules {
			split := strings.Split(line, "{")
			workflows[split[0]] = strings.Split(strings.Trim(split[1], "}"), ",")
		} else {
			split := strings.Split(strings.Trim(line, "{}"), ",")
			newMap := make(map[string]int)
			for _, s := range split {
				split2 := strings.Split(s, "=")
				newMap[split2[0]] = utils.ToInt(split2[1])
			}
			parts = append(parts, newMap)
		}
	}
	fmt.Println(workflows)
	fmt.Println(parts)
	var accepted []map[string]int

	getAcceptance := func(part map[string]int) string {
		curWorkflow := "in"
		//fmt.Print(part, " ")
		for {
			//fmt.Print(curWorkflow, " ")
			if curWorkflow == "A" {
				return "A"
			} else if curWorkflow == "R" {
				return "R"
			}
			workList := workflows[curWorkflow]
			for i, step := range workList {
				if i == len(workList)-1 {
					curWorkflow = step
					break
				}
				wfkey := string(step[0])
				partValue := part[wfkey]
				if step[1] == '<' {
					if partValue < utils.ToInt(strings.Split(step[2:], ":")[0]) {
						curWorkflow = strings.Split(step[2:], ":")[1]
						break
					}
				}
				if step[1] == '>' {
					if partValue > utils.ToInt(strings.Split(step[2:], ":")[0]) {
						curWorkflow = strings.Split(step[2:], ":")[1]
						break
					}
				}
			}
		}
	}

	for _, part := range parts {
		if getAcceptance(part) == "A" {
			accepted = append(accepted, part)
		}
	}
	total := 0
	for _, m := range accepted {
		for _, i := range m {
			total += i
		}
	}
	fmt.Println(total)
	lists := make(map[string][]int)
	rejectMap := make(map[string]map[int]int)
	for _, i := range []string{"x", "a", "m", "s"} {
		rejectMap[i] = make(map[int]int)
	}
	for _, wf := range workflows {
		for _, rule := range wf {
			if strings.Contains(rule, "<") {
				char := string(rule[0])
				lists[char] = append(lists[char], utils.ToInt(strings.Split(rule[2:], ":")[0])-1)
				if strings.Split(rule[2:], ":")[1] == "R" {
					rejectMap[char][utils.ToInt(strings.Split(rule[2:], ":")[0])-1] = -1
				}
			} else if strings.Contains(rule, ">") {
				char := string(rule[0])
				lists[char] = append(lists[char], utils.ToInt(strings.Split(rule[2:], ":")[0]))
				if strings.Split(rule[2:], ":")[1] == "R" {
					rejectMap[char][utils.ToInt(strings.Split(rule[2:], ":")[0])] = 1
				}
			}
		}
	}
	totalVariations := 1
	for s, _ := range lists {
		lists[s] = append(lists[s], 0)
		lists[s] = append(lists[s], 4000)
		slices.Sort(lists[s])
		totalVariations *= len(lists[s])
	}
	fmt.Println(lists)
	fmt.Println("Variants:", totalVariations)

	var allNodes []node
	allNodes = append(allNodes, node{
		curWF:  "in",
		ranges: map[string][2]int{"x": {0, 4000}, "a": {0, 4000}, "s": {0, 4000}, "m": {0, 4000}},
	})

	totalAmount := 0

	for len(allNodes) > 0 {
		var newAllNodes []node
		fmt.Println(len(allNodes))
		for _, curNode := range allNodes {
			if curNode.curWF == "A" {
				amount := 1
				for _, ints := range curNode.ranges {
					amount *= ints[1] - ints[0]
				}
				totalAmount += amount
			} else if curNode.curWF == "R" {
				continue
			}
			nextWorkFlow := workflows[curNode.curWF]
			for _, rule := range nextWorkFlow {
				if !strings.Contains(rule, ":") {
					curNode.curWF = rule
					newAllNodes = append(newAllNodes, curNode)
					continue
				}
				cond := string(rule[0])
				condNum := utils.ToInt(strings.Split(rule[2:], ":")[0])
				nextStep := strings.Split(rule[2:], ":")[1]
				newNode := node{
					curNode.curWF, copyString2IntMap(curNode.ranges),
				}

				low, high := curNode.ranges[cond][0], curNode.ranges[cond][1]
				if strings.Contains(rule, "<") {
					if high < condNum {
						newNode.curWF = nextStep
					} else if low >= condNum {
						continue
					} else {
						lowerRange := [2]int{low, condNum - 1}
						upperRange := [2]int{condNum - 1, high}
						curNode.ranges[cond] = upperRange
						newNode.ranges[cond] = lowerRange
					}
				}
				if strings.Contains(rule, ">") {
					if low > condNum {
						newNode.curWF = nextStep
					} else if high <= condNum {
						continue
					} else {
						lowerRange := [2]int{low, condNum}
						upperRange := [2]int{condNum, high}
						curNode.ranges[cond] = lowerRange
						newNode.ranges[cond] = upperRange
					}

				}
				newNode.curWF = nextStep
				newAllNodes = append(newAllNodes, newNode)
			}
		}
		allNodes = newAllNodes
	}

	fmt.Println(totalAmount)
	fmt.Println(167409079868000)
	fmt.Println(math.MaxInt)
}
