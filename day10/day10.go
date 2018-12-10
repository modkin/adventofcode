package day10

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type star struct {
	x  int
	y  int
	xv int
	yv int
}

func (s *star) move(times int) {
	for i := 0; i < times; i++ {
		s.x += s.xv
		s.y += s.yv
	}
}

func toInt(str string) int {
	ret, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		panic(err)
	}
	return ret
}

func getVolume(stars []star) (vol, xmin, xmax, ymin, ymax int) {
	xmin, xmax, ymin, ymax = stars[0].x, stars[0].x, stars[0].y, stars[0].y
	for _, elem := range stars {
		if elem.x < xmin {
			xmin = elem.x
		}
		if elem.x > xmax {
			xmax = elem.x
		}
		if elem.y < ymin {
			ymin = elem.y
		}
		if elem.y > ymax {
			ymax = elem.y
		}
	}
	vol = (xmax - xmin) * (ymax - ymin)
	return
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func findSmallesBoundingBox(stars []star) int {
	currentVolume, _, _, _, _ := getVolume(stars)
	counter := 0
	for {
		newVolume, _, _, _, _ := getVolume(stars)
		if newVolume > currentVolume {
			return counter
		}
		for idx, _ := range stars {
			stars[idx].move(1)
		}
		currentVolume = newVolume
		counter++
	}
}

func printStars(stars []star, steps int) {
	for idx, _ := range stars {
		stars[idx].move(steps)
	}
	_, xmin, xmax, ymin, ymax := getVolume(stars)
	starPic := make([][]bool, intAbs(ymax-ymin)+1)
	for i := range starPic {
		starPic[i] = make([]bool, intAbs(xmax-xmin)+1)
	}
	for _, elem := range stars {
		//starPic[elem.y+intAbs(ymin)][elem.x+intAbs(xmin)] = true
		starPic[elem.y-ymin][elem.x-xmin] = true
	}
	for y := 0; y <= ymax-ymin; y++ {
		for x := 0; x <= xmax-xmin; x++ {
			if starPic[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}

}

func createStarSlice() []star {
	file, err := os.Open("day10/day10-input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var stars []star
	re := regexp.MustCompile("position=<\\s*([-\\d]*),\\s*([-\\d]*)> velocity=<\\s*([-\\d]*),\\s*([-\\d]*)>")
	for scanner.Scan() {
		word := scanner.Text()
		result := re.FindAllStringSubmatch(word, -1)
		s := star{toInt(result[0][1]),
			toInt(result[0][2]), toInt(result[0][3]),
			toInt(result[0][4])}
		stars = append(stars, s)
	}
	return stars
}

func Task1() int {
	stars := createStarSlice()
	steps := findSmallesBoundingBox(stars)

	stars = createStarSlice()
	fmt.Println(steps - 1)
	printStars(stars, steps-1)

	return 1
}
