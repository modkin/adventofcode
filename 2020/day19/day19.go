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
	var isChar = regexp.MustCompile("[a-z]")
	zeroRule := strings.Split(rules[0], " ")
	substitude := func() bool {
		newZeroRule := make([]string, 0)
		changes := false
		for _, elem := range zeroRule {
			if isChar.MatchString(elem) {
				newZeroRule = append(newZeroRule, elem)
				continue
			} else if elem == "(" || elem == ")" || elem == "|" {
				newZeroRule = append(newZeroRule, elem)
				continue
			}
			changes = true
			tmp := rules[utils.ToInt(elem)]
			tmpSplit := strings.Split(tmp, " ")
			if strings.Contains(tmp, "|") {
				var newRule []string
				if len(tmpSplit) == 5 {
					newRule = []string{"(", tmpSplit[0], tmpSplit[1], "|", tmpSplit[3], tmpSplit[4], ")"}
				} else {
					newRule = []string{"(", tmpSplit[0], "|", tmpSplit[2], ")"}
				}
				newZeroRule = append(newZeroRule, newRule...)
			} else {
				newZeroRule = append(newZeroRule, tmpSplit...)
			}
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
	regex := regexp.MustCompile(regexStr)
	count := 0
	for _, elem := range messages {
		if regex.MatchString(elem) {
			count++
		}
	}
	fmt.Println(count)
}
