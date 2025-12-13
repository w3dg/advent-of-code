package main

import "fmt"

type Graph struct {
	Vertices map[string]*Node
}

type Node struct {
	Val   string
	edges map[string]*Node
	// edges map[string]*Edge
}

// dont need a weighted
// type Edge struct {
// 	weight int
// 	dest   *Node
// }

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]*Node),
	}
}

func (g *Graph) addVertex(name string) {
	if _, ok := g.Vertices[name]; !ok {
		g.Vertices[name] = &Node{
			Val:   name,
			edges: make(map[string]*Node),
		}
	}
}

func (g *Graph) addOrUpdateEdges(nodename string, adj []string) {
	n, ok := g.Vertices[nodename]
	if !ok {
		panic("No such vertex as " + nodename + " to add or update edges to")
	}

	for _, a := range adj {
		_, ok := g.Vertices[a]
		if !ok {
			// fmt.Println(a + " was not found, adding it")
			g.addVertex(a)
		}

		ad, _ := g.Vertices[a]
		n.edges[a] = ad
	}
}

func (g *Graph) String() string {
	s := ""
	for _, node := range g.Vertices {
		s += fmt.Sprintln("Node: ", node)
	}
	return s
}

func (n Node) String() string {
	s := fmt.Sprint(n.Val + " {")
	for e, _ := range n.edges {
		s += fmt.Sprintf(" %v ", e)
	}
	s += " }"
	return s
}
