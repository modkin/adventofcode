package main

import (
	"adventofcode/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"time"
)

type member struct {
	Name    string
	timings []int
}

func genEmptyTimings(length int) []int {
	timings := make([]int, length)
	for i := 0; i < length; i++ {
		timings[i] = math.MaxInt32
	}
	return timings
}

func printDayRanking(allMember []member) {
	for i, mem := range allMember {
		fmt.Println(i+1, " ", mem.Name)
	}
}

func main() {
	jsonByte, _ := ioutil.ReadFile("dailyRanking/678703.json")

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonByte), &data)

	var allMember []member

	//fmt.Println(data["members"].(map[string]interface{})["428833"])

	//for _, elem := range data["members"].(map[string]interface{}) {
	//	name := elem.(map[string]interface{})["name"]
	//	tmp := member{name.(string), genEmptyTimings()}
	//	allMember = append(allMember, tmp)
	//}
	timestamp := time.Now()
	daysDone := timestamp.Day()
	for _, elem := range data["members"].(map[string]interface{}) {
		name := elem.(map[string]interface{})["name"]
		tmp := member{name.(string), genEmptyTimings(daysDone * 2)}

		for i, task := range elem.(map[string]interface{})["completion_day_level"].(map[string]interface{}) {
			tmp.timings[(utils.ToInt(i)-1)*2] = utils.ToInt(task.(map[string]interface{})["1"].(map[string]interface{})["get_star_ts"].(string))
			if val, ok := task.(map[string]interface{})["2"].(map[string]interface{}); ok {
				tmp.timings[(utils.ToInt(i)-1)*2+1] = utils.ToInt(val["get_star_ts"].(string))
			}
		}
		allMember = append(allMember, tmp)

	}

	memberPoints := make([]member, len(allMember))
	copy(memberPoints, allMember)

	//fmt.Println(allMember)
	for k := 0; k < daysDone*2; k++ {
		sort.SliceStable(allMember, func(i, j int) bool { return allMember[i].timings[k] < allMember[j].timings[k] })
		sort.SliceStable(memberPoints, func(i, j int) bool { return memberPoints[i].timings[k] < memberPoints[j].timings[k] })
		for points, _ := range memberPoints {
			if memberPoints[points].timings[k] != math.MaxInt32 {
				memberPoints[points].timings[k] = len(memberPoints) - points
			} else {
				memberPoints[points].timings[k] = 0
			}
		}
	}
	fmt.Println(memberPoints)
	fmt.Printf("%-20v|", "Name")
	for i := 0; i < daysDone*2; i++ {
		fmt.Printf("Day %2v.%v|", (i/2)+1, i%2+1)
	}
	fmt.Print(" Sum\n")
	for _, mem := range memberPoints {
		fmt.Printf("%-20v|", mem.Name)
		for _, point := range mem.timings {
			fmt.Printf(" %6v |", point)
		}
		fmt.Print(" ", utils.SumSlice(mem.timings), "\n")
	}

}
