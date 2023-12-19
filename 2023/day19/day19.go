package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	//partList:
	for _, part := range parts {
		curWorkflow := "in"
		fmt.Print(part, " ")
		for {
			fmt.Print(curWorkflow, " ")
			if curWorkflow == "A" {
				accepted = append(accepted, part)
				break
			} else if curWorkflow == "R" {
				break
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
		fmt.Println()
	}
	total := 0
	for _, m := range accepted {
		for _, i := range m {
			total += i
		}
	}
	fmt.Println(total)
}
