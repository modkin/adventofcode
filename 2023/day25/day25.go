package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func contract(g graph.Graph[string, string], edge graph.Edge[string]) {
	var err error
	v1 := edge.Source
	v2 := edge.Target
	newVertex := v1 + v2
	_ = g.AddVertex(newVertex)
	_ = g.RemoveEdge(v1, v2)
	adjMap, _ := g.AdjacencyMap()
	for _, v := range []string{v1, v2} {
		for s, e := range adjMap[v] {
			edge, err = g.Edge(s, newVertex)
			if err != nil {
				_ = g.AddEdge(s, newVertex, graph.EdgeAttribute("label", "0"))
				edge, _ = g.Edge(s, newVertex)
			}
			edge.Properties.Attributes["label"] = strconv.Itoa(utils.ToInt(edge.Properties.Attributes["label"]) + utils.ToInt(e.Properties.Attributes["label"]))
			err = g.RemoveEdge(e.Source, e.Target)
			if err != nil {
				panic(err)
			}
		}
		err = g.RemoveVertex(v)
		if err != nil {
			panic(err)
		}
	}
}

func algo(g graph.Graph[string, string]) (int, int) {
	// Generate a random number between 0 and 100
	edges, _ := g.Edges()
	for len(edges) != 1 {
		randomNum := rand.Intn(len(edges))
		contract(g, edges[randomNum])
		edges, _ = g.Edges()
	}
	ret := (len(edges[0].Source) / 3) * (len(edges[0].Target) / 3)
	//fmt.Println("num:", edges[0].Properties.Attributes["label"], ret)
	return utils.ToInt(edges[0].Properties.Attributes["label"]), ret
}

func main() {
	file, err := os.Open("2023/day25/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	g := graph.New(graph.StringHash)
	var anyVertex string
	for scanner.Scan() {
		//lines = append(lines, scanner.Text())
		split := strings.Split(scanner.Text(), ": ")
		anyVertex = split[0]
		_, err = g.Vertex(split[0])
		if err != nil {
			_ = g.AddVertex(split[0])
		}
		endNodes := strings.Split(split[1], " ")
		for _, node := range endNodes {
			_, err = g.Vertex(node)
			if err != nil {
				_ = g.AddVertex(node)
			}
			_ = g.AddEdge(split[0], node, graph.EdgeAttribute("label", "1"))
		}

	}

	reachFromVertex := func(vertex string) int {
		counter := 0
		_ = graph.BFS(g, vertex, func(value string) bool {
			counter++
			return false
		})
		return counter
	}

	file2, _ := os.Create("2023/day25/mygraph-orig.gv")
	_ = draw.DOT(g, file2, draw.GraphAttribute("label", "my-graph"))

	file3, _ := os.Create("2023/day25/mygraph.gv")
	_ = draw.DOT(g, file3, draw.GraphAttribute("label", "my-graph"))
	minCut := math.MaxInt
	result := 0
	for i := 0; i < 1000; i++ {
		clone, _ := g.Clone()
		mc, res := algo(clone)
		if mc < minCut {
			minCut = mc
			result = res
			fmt.Println("new min:", minCut, result)
		}
		if i%10 == 0 {
			fmt.Println(result)
		}
	}
	fmt.Println("foo", result)

	allEdges, _ := g.Edges()

outer:
	for i1, edge1 := range allEdges {
		for i2, edge2 := range allEdges {
			for i3, edge3 := range allEdges {
				if i1 < i2 && i2 < i3 {
					_ = g.RemoveEdge(edge1.Source, edge1.Target)
					_ = g.RemoveEdge(edge2.Source, edge2.Target)
					_ = g.RemoveEdge(edge3.Source, edge3.Target)
					if count, _ := g.Order(); reachFromVertex(anyVertex) != count {
						fmt.Println(reachFromVertex(edge1.Source) * reachFromVertex(edge1.Target))
						break outer
					}
					_ = g.AddEdge(edge1.Source, edge1.Target)
					_ = g.AddEdge(edge2.Source, edge2.Target)
					_ = g.AddEdge(edge3.Source, edge3.Target)
				}

			}
		}
	}
}
