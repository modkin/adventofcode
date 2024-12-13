package utils

func AddPoints(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

type Grid2D struct {
	points map[[2]int]string
	xMax   int
	yMax   int
}

func (g Grid2D) Neighbours4Pt(p [2]int) map[[2]int]string {
	ret := make(map[[2]int]string)
	for _, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		tmp := AddPoints(p, offset)
		if g.isInside(tmp) {
			ret[tmp] = g.points[tmp]
		}
	}
	return ret
}

func (g Grid2D) isInside(p [2]int) bool {
	if p[0] < 0 || p[1] < 0 || p[0] > g.xMax || p[1] > g.yMax {
		return false
	}
	return true
}
