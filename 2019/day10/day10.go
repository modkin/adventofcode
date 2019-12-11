package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

func printStars(starmap map[[2]int]int, maxX int, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			tmp := [2]int{x, y}
			if _, ok := starmap[tmp]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func printCount(starmap map[[2]int]int, maxX int, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			tmp := [2]int{x, y}
			if count, ok := starmap[tmp]; ok {
				fmt.Print(count)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func vectorLength(vec [2]int) float64 {
	return math.Sqrt(float64(vec[0]*vec[0]) + float64(vec[1]*vec[1]))
}

func isVisible(star [2]int, otherStar [2]int, starmap map[[2]int]int) bool {
	visible := true
	direction := [2]int{otherStar[0] - star[0], otherStar[1] - star[1]}
	biggerIdx := utils.MaxIdx(direction)
	xstep := float64(direction[0]) / float64(direction[biggerIdx])
	ystep := float64(direction[1]) / float64(direction[biggerIdx])

	for i := utils.IntAbs(direction[biggerIdx]) - 1; i > 0; i-- {
		xDir := xstep * float64(i*direction[biggerIdx]/utils.IntAbs(direction[biggerIdx]))
		yDir := ystep * float64(i*direction[biggerIdx]/utils.IntAbs(direction[biggerIdx]))
		if xDir == float64(int(xDir)) && yDir == float64(int(yDir)) {
			tocheck := [2]int{star[0] + int(xDir), star[1] + int(yDir)}
			if _, ok := starmap[tocheck]; ok {
				visible = false
				break
			}
		}
	}
	return visible
}

func main() {
	filename := "./input"
	content, err := ioutil.ReadFile(filename)
	ylist := strings.Split(string(content), "\n")
	xlist := strings.Split(ylist[0], "")
	xmax, ymax := len(xlist), len(ylist)

	fmt.Println(xmax, " ", ymax)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	starmap := make(map[[2]int]int)
	ycoord := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for xcoord, elem := range line {
			if elem == "#" {
				star := [2]int{xcoord, ycoord}
				starmap[star] = 0
			}
		}
		ycoord++
	}
	for star, _ := range starmap {
		for otherStar, _ := range starmap {
			if otherStar == star {
				continue
			}
			if isVisible(star, otherStar, starmap) {
				starmap[star] = starmap[star] + 1
			}
		}
	}

	//fmt.Println(starmap)
	//printStars(starmap, xmax, ymax)
	//printCount(starmap, xmax, ymax)
	var station [2]int
	count := 0
	for star, starCount := range starmap {
		if starCount > count {
			count = starCount
			station = [2]int{star[0], star[1]}
		}
	}
	if count != 278 {
		fmt.Println("Regression Task 10.1")
	}
	fmt.Println("Station: ", station)
	fmt.Println("Task 10.1: ", count)

	count = 0
	//for true {
	//	for x := station[0]; x > 0; x-- {
	//
	//	}
	//}
	sortedStars := make(map[float64][][2]int)
	for star, _ := range starmap {
		if star == station {
			continue
		}
		direction := [2]int{star[0] - station[0], star[1] - station[1]}
		length := math.Sqrt(float64(direction[0]*direction[0]) + float64(direction[1]*direction[1]))
		gegenkathete := float64(star[0] - station[0])
		angle := math.Asin(gegenkathete / length)
		fmt.Println(star, " ", angle*(180/math.Pi))
		if direction[1] < 0 {
			angle = math.Mod(2*math.Pi+angle, 2*math.Pi)
		} else {
			angle = math.Mod(math.Pi-angle, 2*math.Pi)
		}
		fmt.Println(star, " ", angle*(180/math.Pi))
		sortedStars[angle] = append(sortedStars[angle], [2]int{star[0], star[1]})
		//fmt.Println(star, " ", angle*(180/math.Pi))

	}

	for angle, stars := range sortedStars {
		sort.SliceStable(stars, func(i, j int) bool {
			direction1 := [2]int{stars[i][0] - station[0], stars[i][1] - station[1]}
			direction2 := [2]int{stars[j][0] - station[0], stars[j][1] - station[1]}
			return vectorLength(direction1) < vectorLength(direction2)
		})
		sortedStars[angle] = stars
	}
	sortedAngles := make([]float64, 0, len(sortedStars))
	for angle := range sortedStars {
		sortedAngles = append(sortedAngles, angle)
	}
	sort.Float64s(sortedAngles)
	counter := 0
	for i := 0; i < 205; i++ {
		angle := sortedAngles[counter]
		fmt.Println(i+1, ".: ", sortedStars[angle][0], " ", int(angle*(180/math.Pi)))
		sortedStars[angle] = sortedStars[angle][1:]
		if len(sortedStars[angle]) == 0 {
			delete(sortedStars, angle)
			if counter == len(sortedAngles)-1 {
				sortedAngles = sortedAngles[:counter]
			} else {
				sortedAngles = append(sortedAngles[:counter], sortedAngles[counter+1:]...)
			}
			counter--
		}
		if counter == len(sortedAngles)-1 {
			fmt.Println("reset")
			counter = 0
		} else {
			counter++
		}

	}

}
