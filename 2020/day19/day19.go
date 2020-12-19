package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day19/input"))
	//var regex string
	rules := make(map[int]string)
	messages := make([]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		if len(line) == 2 {
			rules[utils.ToInt(line[0])] = strings.TrimSpace(line[1])
		} else {
			messages = append(messages, line[0])
		}
	}
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"

	var isChar = regexp.MustCompile("[a-z]")
	zeroRule := strings.Split(rules[0], " ")
	counter := 0
	substitude := func() bool {
		if counter == 100 {
			return false
		} else {
			counter++
		}
		newZeroRule := make([]string, 0)
		changes := false
		for _, elem := range zeroRule {
			if isChar.MatchString(elem) {
				newZeroRule = append(newZeroRule, elem)
				continue
			} else if elem == "(" || elem == ")" || elem == "|" || elem == ")*" {
				newZeroRule = append(newZeroRule, elem)
				continue
			}
			changes = true
			tmp := rules[utils.ToInt(elem)]
			tmpSplit := strings.Split(tmp, " ")
			var newRule = tmpSplit
			if strings.Contains(tmp, "|") {
				newRule = []string{"("}
				for _, elem2 := range tmpSplit {
					if elem2 == "|" {
						newRule = append(newRule, "|")
					} else {
						newRule = append(newRule, elem2)
					}
				}
				newRule = append(newRule, ")")
			}
			newZeroRule = append(newZeroRule, newRule...)
		}
		zeroRule = newZeroRule
		return changes
	}

	for substitude() {
		//fmt.Println(zeroRule)
	}

	newZeroRule := []string{"^"}
	for _, elem := range zeroRule {
		newZeroRule = append(newZeroRule, strings.Trim(elem, "\""))
	}
	newZeroRule = append(newZeroRule, "$")
	regexStr := strings.Join(newZeroRule, "")
	//fmt.Println(regexStr)
	regex := regexp.MustCompile(regexStr)
	count := 0
	for _, elem := range messages {
		if regex.MatchString(elem) {
			count++
		}
	}
	fmt.Println(count)
}
