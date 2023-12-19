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

func main() {
	file, err := os.Open("2023/day19/testinput")
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
	rejectMap := make(map[string]map[int]bool)
	for _, i := range []string{"x", "a", "m", "s"} {
		rejectMap[i] = make(map[int]bool)
	}
	for _, wf := range workflows {
		for _, rule := range wf {
			if strings.Contains(rule, "<") {
				char := string(rule[0])
				lists[char] = append(lists[char], utils.ToInt(strings.Split(rule[2:], ":")[0])-1)
				//rejectMap[char][utils.ToInt(strings.Split(rule[2:], ":")[0])-1] = true
			} else if strings.Contains(rule, ">") {
				char := string(rule[0])
				lists[char] = append(lists[char], utils.ToInt(strings.Split(rule[2:], ":")[0]))
				//rejectMap[char][utils.ToInt(strings.Split(rule[2:], ":")[0])] = true
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
	var partList []map[string]int
	var amountList []int

	for ix, x := range lists["x"] {
		if ix == 0 {
			continue
		}
		if _, ok := rejectMap["x"][x]; ok {
			continue
		}
		amountX := (x - lists["x"][ix-1])
		for im, m := range lists["m"] {
			if im == 0 {
				continue
			}
			if _, ok := rejectMap["m"][m]; ok {
				continue
			}
			amountM := amountX * (m - lists["m"][im-1])
			for ia, a := range lists["a"] {
				if ia == 0 {
					continue
				}
				if _, ok := rejectMap["a"][a]; ok {
					continue
				}
				amountA := amountM * (a - lists["a"][ia-1])
				for is, s := range lists["s"] {
					if is == 0 {
						continue
					}
					if _, ok := rejectMap["s"][s]; ok {
						continue
					}
					amount := amountA * (s - lists["s"][is-1])
					amountList = append(amountList, amount)
					newMap := map[string]int{"x": x, "m": m, "a": a, "s": s}
					partList = append(partList)
					partList = append(partList, newMap)
				}
			}
		}
		fmt.Println(len(partList))
	}
	totalAmount := 0
	for i, pl := range partList {
		fmt.Println(len(partList) - i)
		if getAcceptance(pl) == "A" {
			totalAmount += amountList[i]
		}
	}

	fmt.Println(totalAmount)
	fmt.Println(167409079868000)
	fmt.Println(math.MaxInt)
}
