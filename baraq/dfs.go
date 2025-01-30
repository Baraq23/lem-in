package main

import (
	"fmt"
	"strconv"
)

// Point represents a node in the graph with (x, y) coordinates
type Point struct {
	X, Y int
}

// Graph represents the network of rooms
type Graph struct {
	Points map[string]Point    // Stores room names and their coordinates
	Edges  map[string][]string // Adjacency list for connections
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{
		Points: make(map[string]Point),
		Edges:  make(map[string][]string),
	}
}

// AddPoint adds a room to the graph
func (g *Graph) AddPoint(name string, x, y int) {
	g.Points[name] = Point{X: x, Y: y}
}

// AddEdge adds a connection (undirected)
func (g *Graph) AddEdge(a, b string) {
	g.Edges[a] = append(g.Edges[a], b)
	g.Edges[b] = append(g.Edges[b], a)
}

// DFS to find all paths from start to end
func (g *Graph) FindPaths(start, end string, visited map[string]bool, path []string, allPaths *[][]string) {
	// Mark current node as visited
	visited[start] = true
	path = append(path, start)

	// If we reach the end, save the path
	if start == end {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*allPaths = append(*allPaths, pathCopy)
	} else {
		// Explore all unvisited neighbors
		for _, neighbor := range g.Edges[start] {
			if !visited[neighbor] {
				g.FindPaths(neighbor, end, visited, path, allPaths)
			}
		}
	}

	// Backtrack
	visited[start] = false
}

func DFS(connections, points [][]string, startRoom, endRoom []string) ([][]string, error) {
	// Initialize graph
	graph := NewGraph()

	// Add points (rooms)
	for _, n := range points {
		xAxis, yAxis := 0, 0
		var err error
		xAxis, err = strconv.Atoi(n[1])
		if err != nil {
			return nil, fmt.Errorf("ERROR: invalid data format")
		}
		yAxis, err = strconv.Atoi(n[2])
		if err != nil {
			return nil, fmt.Errorf("ERROR: invalid data format")
		}
		graph.AddPoint(string(n[0]), xAxis, yAxis)

	}
	
	// Add connections (edges)
	for _, n := range connections {
		graph.AddEdge(n[0], n[1])

	}
	
	// Find all paths from "O" (Start) to "E" (End)
	visited := make(map[string]bool)
	var allPaths [][]string

	graph.FindPaths(startRoom[0], endRoom[0], visited, []string{}, &allPaths)

	return allPaths, nil
}
