package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type blueprint struct {
	recipies  map[string]map[string]int
	robots    map[string]int
	resources map[string]int
}

func buildifPossible(blue blueprint, target string) bool {
	if target == "obsidian" {
		if blue.resources["ore"] >= blue.recipies[target]["ore"] && blue.resources["clay"] >= blue.recipies[target]["clay"] {
			blue.resources["ore"] -= blue.recipies[target]["ore"]
			blue.resources["clay"] -= blue.recipies[target]["clay"]
			blue.robots[target]++
			if target == "geode" {
				fmt.Println("Geode:", blue.robots["geode"])
			}
			return true

		} else {
			return false
		}
	}
	if target == "geode" {
		if blue.resources["ore"] >= blue.recipies[target]["ore"] && blue.resources["obsidian"] >= blue.recipies[target]["obsidian"] {
			blue.resources["ore"] -= blue.recipies[target]["ore"]
			blue.resources["obsidian"] -= blue.recipies[target]["obsidian"]
			blue.robots[target]++
			//fmt.Println("Geode:", blue.robots["geode"])
			return true

		} else {
			return false
		}
	}
	if target == "clay" || target == "ore" {

		if blue.resources["ore"] >= blue.recipies[target]["ore"] {
			blue.resources["ore"] -= blue.recipies[target]["ore"]
			blue.robots[target]++
			return true
		} else {
			return false
		}
	}
	fmt.Println("Wrong resource")
	return false
}

func getCheckSum(blue blueprint, time int) string {
	ret := ""
	for _, toBuild := range []string{"ore", "clay", "obsidian", "geode"} {
		ret += "-" + strconv.Itoa(blue.robots[toBuild]) //+ "-" + strconv.Itoa(blue.resources[toBuild])
	}
	//ret += "-" + strconv.Itoa(time)
	return ret
}

func main() {
	file, err := os.Open("2022/day19/testinput")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	blueprints := make([]blueprint, 0)

	var newBlue blueprint
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Blueprint") {
			newBlue = blueprint{
				recipies:  make(map[string]map[string]int),
				robots:    map[string]int{"ore": 1},
				resources: map[string]int{"ore": 0, "clay": 0, "obsidian": 0, "geode": 0},
			}
			for scanner.Scan() {
				if scanner.Text() == "" {
					blueprints = append(blueprints, newBlue)
					break
				}
				line := strings.Split(scanner.Text()[2:], " ")
				target := line[1]
				require := make(map[string]int)
				require[strings.Trim(line[5], ".")] = utils.ToInt(line[4])
				if len(line) > 6 {
					require[strings.Trim(line[8], ".")] = utils.ToInt(line[7])
				}
				newBlue.recipies[target] = require
			}
		}
	}
	blueprints = append(blueprints, newBlue)
	//fmt.Println(blueprints)

	//blue := blueprints[0]
	//for time, tobuild := range []string{"", "", "clay", "", "clay", "", "clay", "", "", "", "obsidian", "clay", "", "", "obsidian", "", "", "geode", ""} {
	//	for robo, amount := range blue.robots {
	//		blue.resources[robo] += amount
	//	}
	//	buildifPossible(blue, tobuild)
	//	fmt.Println(time, tobuild)
	//}
	//fmt.Println(blue.robots)
	//fmt.Println(blue.resources)
	//return

	quality := 0
	for idx, blue := range blueprints {
		allBuild := make(map[string]bool)
		allVariants := []blueprint{blue}
		for time := 1; time <= 24; time++ {
			fmt.Println(time, len(allVariants))
			//newAllVariants := make([]blueprint, 0)
			for _, variant := range allVariants {
				for robo, amount := range variant.robots {
					variant.resources[robo] += amount
				}
				//for _, toBuild := range []string{"geode", "obsidian", "clay", "ore"} {
				//	for buildifPossible(variant, toBuild) {
				//	}
				//}

				for _, toBuild := range []string{"ore", "clay", "obsidian", "geode"} {
					newVar := blueprint{
						variant.recipies,
						utils.CopyStringIntMap(variant.robots),
						utils.CopyStringIntMap(variant.resources),
					}
					couldBuild := true
					for couldBuild {
						couldBuild = false
						if buildifPossible(newVar, toBuild) {
							couldBuild = true
							if _, ok := allBuild[getCheckSum(newVar, time)]; !ok {
								allVariants = append(allVariants, newVar)
								allBuild[getCheckSum(newVar, time)] = true
								newVar = blueprint{
									newVar.recipies,
									utils.CopyStringIntMap(newVar.robots),
									utils.CopyStringIntMap(newVar.resources),
								}
							}
						}
					}
				}
			}
			//allVariants = newAllVariants
		}
		max := 0
		for _, variant := range allVariants {
			if tmp := variant.resources["geode"]; tmp > max {
				max = tmp
				fmt.Println("Lokal Max:", max)
			}

		}
		quality += (idx + 1) * max
	}
	fmt.Println(quality)
}
