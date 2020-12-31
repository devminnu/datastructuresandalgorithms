package undirectedgraph

import "fmt"

//Graph Graph
type Graph struct {
	adjacencyList map[string]map[string]bool
}

//New new graph
func New() *Graph {
	return &Graph{
		adjacencyList: map[string]map[string]bool{},
	}
}

//AddVertex AddVertex
func (g *Graph) AddVertex(vertex string) {
	if _, ok := g.adjacencyList[vertex]; !ok {
		g.adjacencyList[vertex] = map[string]bool{}
		return
	}
	fmt.Println("Vertex Already Exist")
}

//AddEdge AddEdge
func (g *Graph) AddEdge(v1, v2 string) {
	_, ok1 := g.adjacencyList[v1]
	_, ok2 := g.adjacencyList[v2]
	if ok1 && ok2 {
		g.adjacencyList[v1][v2] = true
		g.adjacencyList[v2][v1] = true
		return
	}
	if !ok1 {
		fmt.Println("v1 does not exist")
	}
	if !ok2 {
		fmt.Println("v2 does not exist")
	}
}

//RemoveVertex RemoveVertex
func (g *Graph) RemoveVertex(v string) {
	for _, m := range g.adjacencyList {
		delete(m, v)
	}
	delete(g.adjacencyList, v)
	fmt.Println("Removed vertex:", v)
}

//RemoveEdge RemoveEdge
func (g *Graph) RemoveEdge(v1, v2 string) {
	_, ok1 := g.adjacencyList[v1]
	_, ok2 := g.adjacencyList[v2]
	if ok1 && ok2 {
		delete(g.adjacencyList[v1], v2)
		delete(g.adjacencyList[v2], v1)
		fmt.Println("Edge Removed:", v1, v2)
		return
	}
	if !ok1 {
		fmt.Println("v1 does not exist")
	}
	if !ok2 {
		fmt.Println("v2 does not exist")
	}
}

//DFS DFS
func (g *Graph) DFS(v string) {

}

//Print print
func (g *Graph) Print() {
	fmt.Println(g.adjacencyList)
}
