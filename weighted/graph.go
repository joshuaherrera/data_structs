package weighted

import "errors"

type Vertex struct {
	id   interface{}
	cost int
	adj  map[interface{}]int //adj[vertex id] = weight
	from interface{}
	h    int //used for A* search as a hardcoded heuristic value
}

type Graph struct {
	vertices map[interface{}]Vertex
	visited  map[interface{}]bool
	v, e     int
}

var (
	ErrVertNotFound = errors.New("vertex not found")
	ErrSelfLoop     = errors.New("self loops not allowed")
	ErrParallelEdge = errors.New("parallel edges not allowed")
)

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) G() map[interface{}]Vertex {
	return g.vertices
}

func (g *Graph) Visited() map[interface{}]bool {
	return g.visited
}

func (g *Graph) AddVertex(v Vertex) {
	if _, ok := g.vertices[v.id]; !ok {
		g.vertices[v.id] = v
		g.v++
	}

}

func (g *Graph) AddEdge(v, w Vertex, weight int) error {
	if v.id == w.id {
		return ErrSelfLoop
	}
	g.AddVertex(v)
	g.AddVertex(w)

	if _, present := g.vertices[v.id].adj[w.id]; present {
		return ErrParallelEdge
	}

	g.vertices[v.id].adj[w.id] = weight
	g.vertices[w.id].adj[v.id] = weight
	g.e++
	return nil

}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[interface{}]Vertex),
		visited:  make(map[interface{}]bool),
	}
}
