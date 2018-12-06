package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type byTime []timestamp

func toInt(str string) int {
	ret, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return ret
}

type timestamp struct {
	time time.Time
	text string
}

func (a byTime) Len() int      { return len(a) }
func (a byTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byTime) Less(i, j int) bool {
	return a[i].time.Before(a[j].time)
}

func findSleepiestGuard(tslist []timestamp) int {
	guardMap := make(map[int]int)
	sleepMap := make(map[int][]int)
	var currentGuard = -1
	var startsleep time.Time
	for _, elem := range tslist {
		splittedText := strings.Split(elem.text, " ")
		if splittedText[0] == "Guard" {
			var guardId, _ = strconv.Atoi(strings.TrimPrefix(splittedText[1], "#"))
			_, exists := guardMap[guardId]
			if exists {
				currentGuard = guardId
			} else {
				guardMap[guardId] = 0
				sleepMap[guardId] = make([]int, 60)
				currentGuard = guardId
			}
		}
		if splittedText[0] == "falls" {
			startsleep = elem.time
		}
		if splittedText[0] == "wakes" {
			guardMap[currentGuard] += int(elem.time.Sub(startsleep).Minutes())
			for i := startsleep.Minute(); i < elem.time.Minute(); i++ {
				sleepMap[currentGuard][i]++
			}
		}
	}

	var longestSleep = -1
	var longestSleepId = -1
	for id, elem := range guardMap {
		if elem > longestSleep {
			longestSleep = elem
			longestSleepId = id
		}
	}

	var longestMinute = -1
	var longestMinuteId = -1
	for idx, val := range sleepMap[longestSleepId] {
		if val > longestMinute {
			longestMinute = val
			longestMinuteId = idx
			fmt.Println(longestMinute)
		}
	}

	return longestSleepId * longestMinuteId
}

func Task1() int {

	file, err := os.Open("day4/day4input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var timestampList []timestamp

	re := regexp.MustCompile(".([0-9]*)-([0-9]*)-([0-9]*) ([0-9]*):([0-9]*). (.*)")
	for scanner.Scan() {
		word := scanner.Text()
		result := re.FindAllStringSubmatch(word, -1)
		t := time.Date(toInt(result[0][1]),
			time.Month(toInt(result[0][2])),
			toInt(result[0][3]),
			toInt(result[0][4]),
			toInt(result[0][5]), 0, 0, time.UTC)
		v := timestamp{t, result[0][6]}
		timestampList = append(timestampList, v)
	}

	sort.Sort(byTime(timestampList))
	//for _, element := range timestampList {
	//	fmt.Println(element.time.UTC(), element.text)
	//}

	return findSleepiestGuard(timestampList)
}
