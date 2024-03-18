package models

import (
	"fmt"
	"strings"
)

// Graph represents a graph data structure via an adjacency list.
type Graph[T comparable] struct {
	adjacencyList map[T]Node[T]
}

// NewGraph constructs a new instance of Graph.
func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		adjacencyList: make(map[T]Node[T]),
	}
}

// NodeKeys extract all node keys from the adjacency list.
func (g *Graph[T]) NodeKeys() []T {
	keys := make([]T, 0, len(g.adjacencyList))
	for k := range g.adjacencyList {
		keys = append(keys, k)
	}

	return keys
}

// AddNodeIfAbsent adds a new node to the adjacency list representation of the graph if needed.
func (g *Graph[T]) AddNodeIfAbsent(node T) Node[T] {
	// Check if node is already added
	n, ok := g.adjacencyList[node]
	if !ok {
		n = make(Node[T])
		g.adjacencyList[node] = n
	}

	return n
}

// AddEdge connects two nodes with an edge
func (g *Graph[T]) AddEdge(from T, to T) {
	// Add from node.
	fromNode := g.AddNodeIfAbsent(from)
	// Add to node.
	_ = g.AddNodeIfAbsent(to)

	// Connect from and to adjacencyList via an edge
	fromNode.connectsTo(to)
}

// ContainsNode checks if a node is part of the graph via its node key.
func (g *Graph[T]) ContainsNode(key T) bool {
	_, ok := g.adjacencyList[key]
	return ok
}

// LongestPath detects the longest path in a DAG.
func (g *Graph[T]) LongestPath() ([]T, error) {
	longestPath := make([]T, 0)

	for _, k := range g.NodeKeys() {
		currentPath, err := g.ResolveDependencies(k)
		if err != nil {
			return nil, err
		}

		if len(currentPath) > len(longestPath) {
			longestPath = currentPath
		}
	}

	return longestPath, nil
}

// ResolveDependencies runs topological sorting to wire all transitive dependencies of DAG.
func (g *Graph[T]) ResolveDependencies(key T) ([]T, error) {
	results := NewOrderedSet[T]()

	// Run topological sorting algorithm to achieve dependency resolution
	err := g.topSort(key, results, nil)
	if err != nil {
		return nil, err
	}

	return results.items, nil
}

func (g *Graph[T]) topSort(key T, results *OrderedSet[T], visited *OrderedSet[T]) error {
	if visited == nil {
		visited = NewOrderedSet[T]()
	}

	added := visited.add(key)
	if !added {
		return detectCyclicDependency(key, visited)
	}

	n := g.adjacencyList[key]
	for _, edge := range n.edges() {
		err := g.topSort(edge, results, visited.copy())
		if err != nil {
			return err
		}
	}

	results.add(key)

	return nil
}

func detectCyclicDependency[T comparable](key T, visited *OrderedSet[T]) error {
	index := visited.indexOf(key)

	cycle := append(visited.items[index:], key)
	cycleOutput := make([]string, len(cycle))

	for i, k := range cycle {
		cycleOutput[i] = fmt.Sprintf("%v", k)
	}

	return fmt.Errorf("cyclic dependency error: %s", strings.Join(cycleOutput, " -> "))
}
