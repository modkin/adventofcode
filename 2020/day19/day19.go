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
	var validID = regexp.MustCompile("[a-z]")
	zeroRule := strings.Split(rules[0], " ")
	substitude := func() bool {
		newZeroRule := make([]string, 0)
		changes := false
		for _, elem := range zeroRule {
			if validID.MatchString(elem) {
				newZeroRule = append(newZeroRule, elem)
				continue
			} else if elem == "[" || elem == "]" || elem == "][" {
				newZeroRule = append(newZeroRule, elem)
				continue
			}
			changes = true
			tmp := rules[utils.ToInt(elem)]
			tmpSplit := strings.Split(tmp, " ")
			if strings.Contains(tmp, "|") {
				newRule := []string{"[", tmpSplit[0], tmpSplit[3], "][", tmpSplit[1], tmpSplit[4], "]"}
				newZeroRule = append(newZeroRule, newRule...)
			} else {
				newZeroRule = append(newZeroRule, tmpSplit...)
			}
		}
		zeroRule = newZeroRule
		return changes
	}
	fmt.Println(zeroRule)
	for substitude() {
		fmt.Println(zeroRule)
	}
	fmt.Println(zeroRule)
	newZeroRule := []string{"^"}
	for _, elem := range zeroRule {
		newZeroRule = append(newZeroRule, strings.Trim(elem, "\""))
	}
	newZeroRule = append(newZeroRule, "$")
	fmt.Println(strings.Join(newZeroRule, ""))

	tmp := "[w"
	tmp += "-z]"

	validID = regexp.MustCompile(tmp)
	validID.MatchString("adam[23]")
}
