package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type blueprint struct {
	oreForOre  int
	oreForClay int
	oreForObs  int
	clayForObs int
	oreForGeo  int
	obsForGeo  int
}

//go:embed input
var s string

var regex = `Blueprint (\d*): Each ore robot costs (\d*) ore. Each clay robot costs (\d*) ore. Each obsidian robot costs (\d*) ore and (\d*) clay. Each geode robot costs (\d*) ore and (\d*) obsidian.`

func genCS(in []int) (out string) {
	for _, i2 := range in {
		out += strconv.Itoa(i2) + "-"
	}
	return
}

func sumToN(n int) int {
	return (n * (n + 1)) / 2
}

func simulateBlueprint(bb blueprint, maxTime int) (maxGeode int) {
	checksums := make(map[string]bool)
	//allOre := bb.oreForOre + bb.oreForClay + bb.oreForObs + bb.oreForGeo
	maxOreCost := utils.SliceMax([]int{bb.oreForOre, bb.oreForClay, bb.oreForObs, bb.oreForGeo})

	//fmt.Println(maxOreRatio, maxClayRatio)

	var step func(time, ore, clay, obsidian, geo, oreRob, clayRob, obsidianRob, geoRob int)
	step = func(time, ore, clay, obsidian, geo, oreRob, clayRob, obsidianRob, geoRob int) {
		checksum := genCS([]int{time, ore, clay, obsidian, geo, oreRob, clayRob, obsidianRob, geoRob})
		if _, ok := checksums[checksum]; ok {
			return
		} else {
			checksums[checksum] = true
		}
		if time == maxTime {
			if geo > maxGeode {
				maxGeode = geo
			}
			return
		}
		if geo+(maxTime-time)*geoRob+sumToN(maxTime-time) <= maxGeode {
			return
		}

		if obsidian >= bb.obsForGeo && ore >= bb.oreForGeo {
			step(time+1, ore+oreRob-bb.oreForGeo, clay+clayRob, obsidian+obsidianRob-bb.obsForGeo, geo+geoRob, oreRob, clayRob, obsidianRob, geoRob+1)
		}
		if clay >= bb.clayForObs && ore >= bb.oreForObs && obsidianRob <= bb.obsForGeo {
			step(time+1, ore+oreRob-bb.oreForObs, clay+clayRob-bb.clayForObs, obsidian+obsidianRob, geo+geoRob, oreRob, clayRob, obsidianRob+1, geoRob)
		}
		if ore >= bb.oreForClay && clayRob <= bb.clayForObs {
			step(time+1, ore+oreRob-bb.oreForClay, clay+clayRob, obsidian+obsidianRob, geo+geoRob, oreRob, clayRob+1, obsidianRob, geoRob)
		}
		if ore >= bb.oreForOre && oreRob <= maxOreCost {
			step(time+1, ore+oreRob-bb.oreForOre, clay+clayRob, obsidian+obsidianRob, geo+geoRob, oreRob+1, clayRob, obsidianRob, geoRob)
		}
		step(time+1, ore+oreRob, clay+clayRob, obsidian+obsidianRob, geo+geoRob, oreRob, clayRob, obsidianRob, geoRob)
	}
	step(0, 0, 0, 0, 0, 1, 0, 0, 0)
	return maxGeode
}

func main() {

	//blueprints := make([][]string, 0)
	blueprints := make(map[int]blueprint)
	for _, line := range strings.Split(s, "\n") {

		re := regexp.MustCompile(regex)
		data := utils.Map(re.FindStringSubmatch(line)[1:], func(in string) int { return utils.ToInt(in) })
		blueprints[data[0]] = blueprint{data[1], data[2], data[3], data[4], data[5], data[6]}
	}

	fmt.Println(blueprints)
	//start := time.Now()
	part1 := 0
	part2 := 1

	for blueNum, b := range blueprints {
		if blueNum <= 3 {
			maxGeo := simulateBlueprint(b, 32)
			part2 *= maxGeo
			//fmt.Println(blueNum, time.Since(start), maxGeo)
		}
	}

	for blueNum, b := range blueprints {
		maxGeo := simulateBlueprint(b, 24)
		part1 += maxGeo * blueNum
		//fmt.Println(blueNum, time.Since(start))
	}
	fmt.Println("Day 19.1:", part1)
	fmt.Println("Day 19.2:", part2)

}
