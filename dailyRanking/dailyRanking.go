package main

import (
	"adventofcode/utils"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
)

type member struct {
	Name    string
	timings []int
	points  []int
}

func genEmptyTimings(length int) []int {
	timings := make([]int, length)
	for i := 0; i < length; i++ {
		timings[i] = math.MaxInt32
	}
	return timings
}

func main() {
	jsonByte, _ := os.ReadFile("dailyRanking/678703.json")

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonByte), &data)
	if err != nil {
		panic(err)
	}

	var allMember []member

	//fmt.Println(data["members"].(map[string]interface{})["428833"])

	//for _, elem := range data["members"].(map[string]interface{}) {
	//	name := elem.(map[string]interface{})["name"]
	//	tmp := member{name.(string), genEmptyTimings()}
	//	allMember = append(allMember, tmp)
	//}
	//timestamp := time.Now()
	daysDone := 15
	for _, elem := range data["members"].(map[string]interface{}) {
		name := elem.(map[string]interface{})["name"]
		tmp := member{name.(string), genEmptyTimings(daysDone * 2), make([]int, daysDone*2)}

		for i, task := range elem.(map[string]interface{})["completion_day_level"].(map[string]interface{}) {
			tmp.timings[(utils.ToInt(i)-1)*2] = int(task.(map[string]interface{})["1"].(map[string]interface{})["get_star_ts"].(float64))
			if val, ok := task.(map[string]interface{})["2"].(map[string]interface{}); ok {
				tmp.timings[(utils.ToInt(i)-1)*2+1] = int(val["get_star_ts"].(float64))
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
		for points := range memberPoints {
			if memberPoints[points].timings[k] != math.MaxInt32 {
				memberPoints[points].points[k] = len(memberPoints) - points
			} else {
				memberPoints[points].points[k] = 0
			}
		}
	}
	//fmt.Println(memberPoints)
	fmt.Printf("%-24v|", "Name")
	for i := 0; i < daysDone*2; i++ {
		fmt.Printf("%2v|", (i/2)+1)
	}
	fmt.Print(" Sum\n")
	for _, mem := range memberPoints {
		fmt.Printf("%-24v|", mem.Name)
		for _, point := range mem.points {
			fmt.Printf("%2v|", point)
		}
		fmt.Print(" ", utils.SumSlice(mem.points), "\n")
	}
	fmt.Println()
	for i := 0; i < daysDone*2; i++ {
		difference := memberPoints[2].timings[i] - memberPoints[1].timings[i]
		//fmt.Println("Day: ", i/2, ".", i%2)
		if utils.IntAbs(difference) < 120 && difference != 0 {
			fmt.Print("Day: ", (i/2)+1, ".", i%2, " ", difference, "\n")
		}
	}
}
