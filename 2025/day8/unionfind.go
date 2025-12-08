package main

import (
	"fmt"
	"strings"
)

type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

// NewUnionFind initializes a new Union-Find instance.
func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
	}
}

// Add inserts a new element into the union-find structure.
func (uf *UnionFind[T]) Add(x T) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x // Set itself as parent
		uf.rank[x] = 0   // Initialize rank
	}
}

// Find returns the root of the component containing the element.
func (uf *UnionFind[T]) Find(x T) T {
	if parent, ok := uf.parent[x]; ok {
		if parent != x {
			uf.parent[x] = uf.Find(parent) // Path compression
		}
		return uf.parent[x]
	}
	// If x is not in the parent map, set it as its own parent
	uf.parent[x] = x
	return x
}

// Union merges two components together.
func (uf *UnionFind[T]) Union(x, y T) {
	uf.Add(x) // Ensure x is added before union
	uf.Add(y) // Ensure y is added before union

	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// If roots are different, combine the trees
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// String provides a string representation of the Union-Find structure.
func (uf *UnionFind[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Union-Find Structure:\n")

	// Create a map to group connected components
	groups := uf.GetGroups()

	// Append each group to the string
	for root, elements := range groups {
		sb.WriteString(fmt.Sprintf("Root %v: %v (%v)\n", root, elements, len(elements)))
	}

	return sb.String()
}

// Returns the connected components of the union find.
// Each entry has the head of the group and a list of the components within that group.
func (uf *UnionFind[T]) GetGroups() map[T][]T {
	groups := make(map[T][]T)

	for element := range uf.parent {
		root := uf.Find(element)
		groups[root] = append(groups[root], element)
	}

	return groups
}

// Returns a map of Size of the group => Number of circuits of that length
func (uf *UnionFind[T]) GetGroupSizes() map[int]int {
	m := make(map[int]int)

	// Create a map to group connected components
	groups := uf.GetGroups()

	// Append each group to the string
	for _, elements := range groups {
		circuitLength := len(elements)

		if v, ok := m[circuitLength]; ok {
			m[circuitLength] = v + 1
		} else {
			m[circuitLength] = 1
		}
	}

	return m
}
