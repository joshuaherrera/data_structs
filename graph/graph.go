package graph

//a simple graph
//need to be able to add edges, vertices, find degree of vertices
//connected to v, be able to list vertices connected to v, and keep
//track of edges and vertices

import (
	"errors"
)

var (
	ErrVertNotFound = errors.New("vertex not found")
	ErrSelfLoop     = errors.New("self loops not allowed")
	ErrParallelEdge = errors.New("parallel edges not allowed")
)

type Graph struct {
	adjList map[interface{}]map[interface{}]struct{}
	v, e    int
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) AddVertex(v interface{}) {
	//check if exists
	list, present := g.adjList[v]
	if !present {
		list = make(map[interface{}]struct{})
		g.adjList[v] = list
		g.v++
	}
}

func (g *Graph) AddEdge(v interface{}, w interface{}) error {
	if v == w {
		return ErrSelfLoop
	}

	g.AddVertex(v)
	g.AddVertex(w)

	//use this form with semicolon to keep ok in this scope only
	if _, present := g.adjList[v][w]; present {
		return ErrParallelEdge
	}
	//look into struct here
	g.adjList[v][w] = struct{}{}
	g.adjList[v][w] = struct{}{}
	g.e++
	return nil

}

func (g *Graph) Degree(v interface{}) (int, error) {
	val, ok := g.adjList[v]
	if !ok {
		return 0, ErrVertNotFound
	}
	return len(val), nil
}

func (g *Graph) Adj(v interface{}) (interface{}, error) {
	deg, err := g.Degree(v)
	if err != nil {
		return nil, ErrVertNotFound
	}

	adj := make([]interface{}, deg)
	i := 0
	for key := range g.adjList[v] {
		adj[i] = key
		i++
	}

	return adj, nil
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[interface{}]map[interface{}]struct{}),
		v:       0,
		e:       0,
	}
}
