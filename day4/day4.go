package day4

import (
	"bufio"
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

func createMaps(tslist []timestamp) (map[int]int, map[int][]int) {
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
	return guardMap, sleepMap
}

func findSleepiestGuard(tslist []timestamp) int {
	guardMap, sleepMap := createMaps(tslist)
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
		}
	}
	return longestSleepId * longestMinuteId
}

func findMaxSleepMinute(tslist []timestamp) int {
	_, sleepMap := createMaps(tslist)
	var maxMinute, minuteId, guardId = -1, -1, -1
	for id, elem := range sleepMap {
		for id2, elem2 := range elem {
			if elem2 > maxMinute {
				maxMinute = elem2
				minuteId = id2
				guardId = id
			}
		}
	}
	return guardId * minuteId

}

func createTimeStampList() []timestamp {
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

	return timestampList
}

func Task1() int {
	return findSleepiestGuard(createTimeStampList())
}
func Task2() int {
	return findMaxSleepMinute(createTimeStampList())
}
