package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Planet struct {
	pos [3]int
	vel [3]int
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var planets []Planet
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		var pos [3]int
		for i, planet := range coords {
			p, _ := strconv.Atoi(strings.TrimSuffix(strings.Split(planet, "=")[1], ">"))
			pos[i] = p
		}
		newPlanet := Planet{
			pos: pos,
			vel: [3]int{},
		}
		planets = append(planets, newPlanet)
	}
	fmt.Println(planets)

	//var states []int64
	timesteps := 1000
	var energy int64
	for time := 0; time < timesteps; time++ {
		///update velocity
		for idx, _ := range planets {
			for _, otherPlanet := range planets {
				if planets[idx] == otherPlanet {
					continue
				}
				for i := 0; i <= 2; i++ {
					if planets[idx].pos[i] > otherPlanet.pos[i] {
						planets[idx].vel[i]--
					} else if planets[idx].pos[i] < otherPlanet.pos[i] {
						planets[idx].vel[i]++
					}
				}
			}
		}
		/// update position
		for idx, _ := range planets {
			for i := 0; i <= 2; i++ {
				planets[idx].pos[i] += planets[idx].vel[i]
			}
		}
	}

	energy = 0
	for _, planet := range planets {
		kin, pot := 0, 0
		for i := 0; i <= 2; i++ {
			kin += utils.IntAbs(planet.pos[i])
			pot += utils.IntAbs(planet.vel[i])
		}
		energy += int64(kin) * int64(pot)
		fmt.Println(energy)
	}

	fmt.Println("Task 12.1: ", energy)
}
