package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type hail struct {
	pos []float64
	dir []float64
}

func getLineIntersection(one hail, two hail) (bool, []float64) {

	// Extrahiere die Parameter der Geraden
	x1, y1 := one.pos[0], one.pos[1]
	dx1, dy1 := one.dir[0], one.dir[1]

	x2, y2 := two.pos[0], two.pos[1]
	dx2, dy2 := two.dir[0], two.dir[1]

	// Überprüfe, ob die Geraden parallel sind
	det := dx1*dy2 - dy1*dx2
	if math.Abs(det) < 1e-9 { // Prüfe auf Parallelität mit einem kleinen Toleranzwert
		return false, []float64{} // Geraden sind parallel, kein Schnittpunkt
	}

	// Berechne den Schnittpunkt
	t := ((x2-x1)*dy2 - (y2-y1)*dx2) / det
	u := ((x2-x1)*dy1 - (y2-y1)*dx1) / det

	// Überprüfe, ob die Geraden sich hinter ihren Startpunkten schneiden
	if t > 0 && u > 0 {
		intersection := []float64{x1 + t*dx1, y1 + t*dy1}
		return true, intersection
	}
	return false, []float64{}
}

func stringToFloatSlice(in []string) []float64 {
	out := make([]float64, len(in))
	for i, s := range in {
		out[i], _ = strconv.ParseFloat(s, 64)
	}
	return out
}

func main() {
	file, err := os.Open("2023/day24/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var hailList []hail
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "@")
		pos := strings.Split(strings.Replace(split[0], " ", "", -1), ",")
		dir := strings.Split(strings.Replace(split[1], " ", "", -1), ",")
		newHail := hail{
			pos: stringToFloatSlice(pos),
			dir: stringToFloatSlice(dir),
		}
		hailList = append(hailList, newHail)
	}

	crossPoints := make(map[[3]float64]bool)
	counter := 0
	for i1, one := range hailList {
		for i2, two := range hailList {
			if i1 < i2 {
				continue
			}
			cross, intersection := getLineIntersection(one, two)
			if cross {
				if intersection[0] >= 200000000000000 && intersection[0] <= 400000000000000 && intersection[1] >= 200000000000000 && intersection[1] <= 400000000000000 {
					//xdirOne := (intersection[0] - one.pos[0]) / math.Abs(intersection[0]-one.pos[0])
					//xdirTwo := (intersection[0] - two.pos[0]) / math.Abs(intersection[0]-two.pos[0])
					//if xdirOne == xdirTwo {
					//fmt.Println(one.pos)
					//fmt.Println(two.pos)
					//fmt.Println(intersection)
					//fmt.Println()
					//interArray := [3]float64{math.Round(intersection[0]*10) / 10, math.Round(intersection[1]*10) / 10}
					interArray := [3]float64{intersection[0], intersection[1]}
					crossPoints[interArray] = true
					counter++
					//}
				}
			}
		}
	}
	fmt.Println("Day 24.1:", len(crossPoints))
	fmt.Println(counter)
	//fmt.Println(crossPoints)
}
