package day13

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"sort"
)

type trainLex []train

type train struct {
	x        int
	y        int
	xspeed   int
	yspeed   int
	position rune
	//0 = left; 1 == straight; 2 == right
	turn int
}

func (t trainLex) Len() int      { return len(t) }
func (t trainLex) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t trainLex) Less(i, j int) bool {
	if t[i].y < t[j].y {
		return true
	}
	if t[i].x < t[j].x {
		return true
	}
	return false
}

func removeTrain(t train, trains []train) []train {
	for idx, train := range trains {
		if train == t {
			return append(trains[:idx], trains[idx+1:]...)
		}
	}
	return []train{}
}

func getTrainSymbol(t train) rune {
	switch t.xspeed {
	case 1:
		return '>'
	case -1:
		return '<'
	}
	switch t.yspeed {
	case 1:
		return 'v'
	case -1:
		return '^'
	}
	return 'E'
}

func (t train) String() string {
	return fmt.Sprintf("x: %d y: %d xs: %d ys: %d p: %c", t.x, t.y, t.xspeed, t.yspeed, t.position)
}

func checkCrash(idx int, trains []train) (bool, int) {
	for i, train := range trains {
		if i != idx {
			if train.x == trains[idx].x && train.y == trains[idx].y {
				return true, i
			}
		}
	}
	return false, -1
}

func printPlan(plan [][]rune) {
	for _, elem := range plan {
		fmt.Println(string(elem))
	}
}

func parsePlan() [][]rune {
	var plan [][]rune
	scanner := bufio.NewScanner(utils.OpenFile("day13/day13-input.txt"))
	for scanner.Scan() {
		line := []rune(scanner.Text())
		plan = append(plan, line)
	}
	return plan
}

func findTrains(plan [][]rune) []train {
	var trains []train
	for y, elem := range plan {
		for x, char := range elem {
			switch char {
			case '>':
				trains = append(trains, train{x, y, 1, 0, '-', 0})
			case '<':
				trains = append(trains, train{x, y, -1, 0, '-', 0})
			case '^':
				trains = append(trains, train{x, y, 0, -1, '|', 0})
			case 'v':
				trains = append(trains, train{x, y, 0, 1, '|', 0})
			}
		}
	}
	return trains
}

func moveTrains(plan [][]rune, trains []train) (bool, []int) {
	sort.Sort(trainLex(trains))
	var crash []int
	for idx, _ := range trains {
		train := &trains[idx]
		plan[train.y][train.x] = train.position
		train.x += train.xspeed
		train.y += train.yspeed
		train.position = plan[train.y][train.x]
		switch plan[train.y][train.x] {
		case '\\':
			train.xspeed, train.yspeed = train.yspeed, train.xspeed
		case '/':
			train.xspeed, train.yspeed = -train.yspeed, -train.xspeed
		case '+':
			switch train.turn {
			case 0:
				train.xspeed, train.yspeed = train.yspeed, -train.xspeed
			case 2:
				train.xspeed, train.yspeed = -train.yspeed, train.xspeed
			}
			train.turn = (train.turn + 1) % 3
		}
		plan[train.y][train.x] = getTrainSymbol(*train)
		crashed, _ := checkCrash(idx, trains)
		if crashed {
			crash = []int{train.x, train.y}
			plan[train.y][train.x] = 'X'
			return true, crash
		}
	}
	return false, crash
}

func Task1() {
	plan := parsePlan()
	trains := findTrains(plan)
	var crashCoord []int
	crash := false
	for !crash {
		crash, crashCoord = moveTrains(plan, trains)
	}
	fmt.Println(crashCoord)
}

func Task2() {
	plan := parsePlan()
	trains := findTrains(plan)
	for len(trains) != 1 {
		//_, _, trains = moveTrains(plan, trains)
	}
	fmt.Println(trains[0].x, trains[0].y)
}
