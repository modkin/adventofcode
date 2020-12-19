package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day19/testinput"))
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
			} else if elem == "(" || elem == ")" || elem == "|" || elem == ")*" {
				newZeroRule = append(newZeroRule, elem)
				continue
			}
			changes = true
			tmp := rules[utils.ToInt(elem)]
			tmpSplit := strings.Split(tmp, " ")
			var newRule = tmpSplit
			if elem == "8" {
				newRule = []string{"(", "42", ")*"}
			} else if elem == "11" {
				//newRule = []string{"(", "42", ")*", "(", "31", ")*"}
				newRule = []string{"("}
				for i := 1; i < 100; i++ {
					tmp1 := make([]string, 0)
					tmp2 := make([]string, 0)
					for count := 1; count <= i; count++ {
						tmp1 = append(tmp1, "42")
						tmp2 = append(tmp2, "31")
					}
					newRule = append(newRule, tmp1...)
					newRule = append(newRule, tmp2...)
					if i != 99 {
						newRule = append(newRule, "|")
					}
				}
				newRule = append(newRule, ")")
			} else if strings.Contains(tmp, "|") {
				if len(tmpSplit) == 5 {
					newRule = []string{"(", tmpSplit[0], tmpSplit[1], "|", tmpSplit[3], tmpSplit[4], ")"}
				} else {
					newRule = []string{"(", tmpSplit[0], "|", tmpSplit[2], ")"}
				}
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
