package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type block struct {
	id       int
	startIdx int
	len      int
}

func main() {

	var files []block
	var freeSpace []block
	id := 0
	idx := 0
	isDisk := true
	dist := make([]int, 0)
	for _, line := range utils.ReadFileIntoLines("2024/day9/input") {
		split := strings.Split(line, "")
		for _, s := range split {
			blockLen := utils.ToInt(s)
			if isDisk {
				for i := 0; i < utils.ToInt(s); i++ {
					dist = append(dist, id)
				}
				files = append(files, block{id, idx, blockLen})
			} else {
				for i := 0; i < utils.ToInt(s); i++ {
					dist = append(dist, -1)
				}
				freeSpace = append(freeSpace, block{id, idx, blockLen})
			}
			idx += blockLen
			if isDisk {
				id++
			}
			isDisk = !isDisk
		}
	}

	id--

	backwards := len(dist) - 1
	for i := 0; i < backwards; i++ {
		if dist[i] == -1 {
			dist[i] = dist[backwards]
			dist[backwards] = -1
			for backwards--; dist[backwards] == -1; backwards-- {
			}
		}
	}

	checksum := 0
	for i, n := range dist {
		if n != -1 {
			checksum += i * n
		}
	}
	fmt.Println("Day 9.1:", checksum)
	for fileId := len(files) - 1; fileId > 0; fileId-- {
		f := files[fileId]
		for i, b := range freeSpace {
			if b.startIdx >= f.startIdx {
				break
			}
			if f.len == b.len {
				files[fileId].startIdx = b.startIdx
				freeSpace = append(freeSpace[:i], freeSpace[i+1:]...)
				break
			} else if f.len < b.len {
				files[fileId].startIdx = b.startIdx
				freeSpace[i].startIdx += f.len
				freeSpace[i].len -= f.len
				break
			}
		}
	}

	checksum = 0
	for _, f := range files {
		for i := 0; i < f.len; i++ {
			//fmt.Println(f.id, f.id*(f.startIdx+i))
			checksum += f.id * (f.startIdx + i)
		}
	}
	fmt.Println("Day 9.2:", checksum)

}
