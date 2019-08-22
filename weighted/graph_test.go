package weighted

import (
	"testing"
)

func TestVertices(t *testing.T) {
	v1 := Vertex{
		id:   1,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v2 := Vertex{
		id:   2,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v3 := Vertex{
		id:   3,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v4 := Vertex{
		id:   4,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}

	g := NewGraph()
	g.AddVertex(v1)
	if g.V() != 1 {
		t.Errorf("Error: got %v vertices, wanted %v", g.V(), 1)
	}
	g.AddVertex(v2)
	if g.V() != 2 {
		t.Errorf("Error: got %v vertices, wanted %v", g.V(), 2)
	}
	g.AddVertex(v3)
	if g.V() != 3 {
		t.Errorf("Error: got %v vertices, wanted %v", g.V(), 3)
	}
	g.AddVertex(v4)
	if g.V() != 4 {
		t.Errorf("Error: got %v vertices, wanted %v", g.V(), 4)
	}

}

func TestEdges(t *testing.T) {
	v1 := Vertex{
		id:   1,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v2 := Vertex{
		id:   2,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v3 := Vertex{
		id:   3,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v4 := Vertex{
		id:   4,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}

	g := NewGraph()
	g.AddEdge(v1, v2, 5)
	if g.E() != 1 {
		t.Errorf("Error: got %v edges, wanted %v", g.E(), 1)
	}
	g.AddEdge(v1, v3, 7)
	if g.E() != 2 {
		t.Errorf("Error: got %v edges, wanted %v", g.E(), 2)
	}
	g.AddEdge(v2, v3, 1)
	if g.E() != 3 {
		t.Errorf("Error: got %v edges, wanted %v", g.E(), 3)
	}
	g.AddEdge(v2, v4, 10)
	if g.E() != 4 {
		t.Errorf("Error: got %v edges, wanted %v", g.E(), 4)
	}
	g.AddEdge(v3, v4, 6)
	if g.E() != 5 {
		t.Errorf("Error: got %v edges, wanted %v", g.E(), 5)
	}

	//fmt.Println(g.G())

}

func TestShortestPath(t *testing.T) {
	v1 := Vertex{
		id:   1,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v2 := Vertex{
		id:   2,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v3 := Vertex{
		id:   3,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v4 := Vertex{
		id:   4,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v5 := Vertex{
		id:   5,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}
	v6 := Vertex{
		id:   6,
		dist: 999999,
		adj:  make(map[interface{}]int),
	}

	g := NewGraph()
	g.AddEdge(v1, v2, 5)
	g.AddEdge(v1, v3, 7)
	g.AddEdge(v2, v3, 1)
	g.AddEdge(v2, v4, 10)
	g.AddEdge(v3, v4, 6)
	g.AddEdge(v5, v6, 777)

	cases := []struct {
		v, w, want int
	}{
		{1, 4, 12},
		{1, 3, 6},
		{2, 4, 7},
		{1, 2, 5},
		{2, 3, 1},
		{1, 6, 999999999},
	}

	for _, c := range cases {
		got := g.ShortestPath(c.v, c.w)
		if got != c.want {
			t.Errorf("Error: got %v for src %v to dest %v, wanted %v", got, c.v, c.w, c.want)
		}
	}

}
