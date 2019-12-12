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

func calcEnergy(planets [4]Planet) (energy int) {
	energy = 0
	for _, planet := range planets {
		kin, pot := 0, 0
		for i := 0; i <= 2; i++ {
			kin += utils.IntAbs(planet.pos[i])
			pot += utils.IntAbs(planet.vel[i])
		}
		energy += kin * pot
	}
	return
}

func main() {
	file, err := os.Open("./testInput1")
	if err != nil {
		panic(err)
	}

	var planets [4]Planet
	scanner := bufio.NewScanner(file)
	planetPos := 0
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
		planets[planetPos] = newPlanet
		planetPos++
	}
	fmt.Println(planets)

	states := make(map[[4]Planet]bool)
	//timesteps := 10 * 4686774924
	//for time := 0; time < timesteps; time++ {
	for time := 0; ; time++ {
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

		if _, ok := states[planets]; ok {
			fmt.Println("time: ", time)
			return
		}
		states[planets] = true
		if time%1000000 == 0 {
			fmt.Println(time)
		}
	}

	fmt.Println(planets)

	fmt.Println("Task 12.1: ", calcEnergy(planets))
}
