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

func (s *star) move() {
	s.x += s.xv
	s.y += s.yv
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

func findSmallesBoundingBox(stars []star) {
	currentVolume, _, _, _, _ := getVolume(stars)
	oldStars := stars
	for {
		newVolume, _, _, _, _ := getVolume(stars)
		if newVolume > currentVolume {
			stars = oldStars
			break
		}
		oldStars = stars
		for idx, _ := range stars {
			stars[idx].move()
		}
		currentVolume = newVolume
		printStars(stars)
	}
}

func printStars(stars []star) {
	_, xmin, xmax, ymin, ymax := getVolume(stars)
	starPic := make([][]rune, intAbs(ymax-ymin)+1)
	for i := range starPic {
		starPic[i] = make([]rune, intAbs(xmax-xmin)+1)
	}
	for _, elem := range stars {
		starPic[elem.y+intAbs(ymin)][elem.x+intAbs(xmin)] = '#'
	}
	for y := 0; y < ymax-ymin; y++ {
		for x := 0; x < xmax-xmin; x++ {
			fmt.Print(string(starPic[y][x]))
		}
		fmt.Println()
	}

}

func createStarSlice() []star {
	file, err := os.Open("day10/day10-example.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var stars []star

	//re := regexp.MustCompile(".*<([0-9]*), ([0-9]*)>.*<([0-9]*), ([0-9]*)>")
	//re := regexp.MustCompile("position=< ?(-?[0-9]*)  ?(-?[0-9])> velocity=< ?(-?[0-9]*)  ?(-?[0-9])>")
	//re := regexp.MustCompile("position=</s*([-/d]*),/s*([-/d])> velocity=</s*([-/d]*),/s*([-/d])>")
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
	printStars(stars)
	findSmallesBoundingBox(stars)
	printStars(stars)
	return 1
}
