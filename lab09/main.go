package main

import (
	"lab09/graph/undirectedgraph"
)

func main() {
	g := undirectedgraph.New()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	// A:B,C  | B:A,E | C:A,D,E | D:C  | E:B,C
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "E")
	g.AddEdge("C", "D")
	g.AddEdge("C", "E")
	g.Print()
	g.RemoveVertex("B")
	g.Print()
	g.RemoveEdge("A", "C")
	g.Print()

}
