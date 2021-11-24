package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type reindeer struct {
	speed     int //in km/s
	power     int //how long before rest
	rest      int //how long to rest
	pos       int
	powerLeft int
	restLeft  int
	points    int
}

func main() {

	file, err := os.Open("2015/day14/input.txt")
	if err != nil {
		panic(err)
	}

	var reindeers []reindeer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		reindeers = append(reindeers, reindeer{
			speed:     utils.ToInt(lineSplit[3]),
			power:     utils.ToInt(lineSplit[6]),
			rest:      utils.ToInt(lineSplit[13]),
			powerLeft: utils.ToInt(lineSplit[6]),
			restLeft:  utils.ToInt(lineSplit[13]),
		})
	}
	seconds := 2503
	maxDistance := 0
	for _, r := range reindeers {
		time := 0
		distance := 0
		for {
			timeLeft := seconds - time
			if timeLeft > r.power {
				time += r.power
				distance += r.speed * r.power
				time += r.rest
				if time >= seconds {
					break
				}
			} else {
				distance += r.speed * timeLeft
				break
			}
		}
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	fmt.Println("Day 14.1:", maxDistance)
	for time := 1; time <= seconds; time++ {
		for idx, r := range reindeers {
			if r.powerLeft > 0 {
				reindeers[idx].pos += r.speed
				reindeers[idx].powerLeft--
			} else if r.restLeft > 0 {
				reindeers[idx].restLeft--
				if reindeers[idx].restLeft == 0 {
					reindeers[idx].powerLeft = r.power
					reindeers[idx].restLeft = r.rest
				}
			}
		}
		leaderIdx := 0
		leaderPos := 0
		for idx, r := range reindeers {
			if r.pos > leaderPos {
				leaderPos = r.pos
				leaderIdx = idx
			}
		}
		reindeers[leaderIdx].points++
	}
	mostPoints := 0
	for _, r := range reindeers {
		if r.points > mostPoints {
			mostPoints = r.points
		}
	}
	fmt.Println("Day 14.2:", mostPoints)
}
