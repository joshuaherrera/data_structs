package weighted

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVertices(t *testing.T) {
	v1 := Vertex{
		id:   1,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v2 := Vertex{
		id:   2,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v3 := Vertex{
		id:   3,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v4 := Vertex{
		id:   4,
		cost: 999999,
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
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v2 := Vertex{
		id:   2,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v3 := Vertex{
		id:   3,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v4 := Vertex{
		id:   4,
		cost: 999999,
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

func TestDijkstraShortestPath(t *testing.T) {
	v1 := Vertex{
		id:   1,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v2 := Vertex{
		id:   2,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v3 := Vertex{
		id:   3,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v4 := Vertex{
		id:   4,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v5 := Vertex{
		id:   5,
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	v6 := Vertex{
		id:   6,
		cost: 999999,
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
		wpath      []interface{}
	}{
		{1, 4, 12, []interface{}{1, 2, 3, 4}},
		{1, 3, 6, []interface{}{1, 2, 3}},
		{2, 4, 7, []interface{}{2, 3, 4}},
		{1, 2, 5, []interface{}{1, 2}},
		{2, 3, 1, []interface{}{2, 3}},
		{1, 6, 999999999, nil},
	}

	for _, c := range cases {
		g.ClearVisited()
		path, got := g.ShortestPath(c.v, c.w)
		if got != c.want {
			t.Errorf("Error: got %v for src %v to dest %v, wanted %v", got, c.v, c.w, c.want)
		}
		if reflect.DeepEqual(path, c.wpath) == false {
			t.Errorf("Error: got %v path, wanted %v path", path, c.wpath)
		}
	}

}

func TestDijkstraTwo(t *testing.T) {
	/*uses same graph as A* to see difference in work performed*/
	s := Vertex{
		id:   "S",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    10,
	}
	a := Vertex{
		id:   "A",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    9,
	}
	b := Vertex{
		id:   "B",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    7,
	}
	c := Vertex{
		id:   "C",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    8,
	}
	d := Vertex{
		id:   "D",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    8,
	}
	e := Vertex{
		id:   "E",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    0,
	}
	f := Vertex{
		id:   "F",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    6,
	}
	g := Vertex{
		id:   "G",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    3,
	}
	h := Vertex{
		id:   "H",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    6,
	}
	i := Vertex{
		id:   "I",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    4,
	}
	j := Vertex{
		id:   "J",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    4,
	}
	k := Vertex{
		id:   "K",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    3,
	}
	l := Vertex{
		id:   "L",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    6,
	}

	graph := NewGraph()
	graph.AddEdge(s, a, 7)
	graph.AddEdge(s, b, 2)
	graph.AddEdge(s, c, 3)
	graph.AddEdge(a, b, 3)
	graph.AddEdge(a, d, 4)
	graph.AddEdge(b, d, 4)
	graph.AddEdge(b, h, 1)
	graph.AddEdge(d, f, 5)
	graph.AddEdge(h, f, 3)
	graph.AddEdge(h, g, 2)
	graph.AddEdge(g, e, 2)
	graph.AddEdge(c, l, 2)
	graph.AddEdge(l, i, 4)
	graph.AddEdge(l, j, 4)
	graph.AddEdge(i, j, 6)
	graph.AddEdge(i, k, 4)
	graph.AddEdge(j, k, 4)
	graph.AddEdge(k, e, 5)

	cases := []struct {
		v, w  string
		want  int
		wpath []interface{}
	}{
		{"S", "E", 7, []interface{}{"S", "B", "H", "G", "E"}},
		/*		{1, 3, 6, []interface{}{1, 2, 3}},
				{2, 4, 7, []interface{}{2, 3, 4}},
				{1, 2, 5, []interface{}{1, 2}},
				{2, 3, 1, []interface{}{2, 3}},
				{1, 6, 999999999, nil},*/
	}
	for _, c := range cases {
		graph.ClearVisited()
		path, got := graph.ShortestPath(c.v, c.w)
		if got != c.want {
			t.Errorf("Error: got %v cost, wanted %v", got, c.want)
		}
		if reflect.DeepEqual(path, c.wpath) == false {
			t.Errorf("Error: got %v path, wanted %v", path, c.wpath)
		}
	}
}

func TestAStarShortestPath(t *testing.T) {
	/*used https://www.youtube.com/watch?v=ySN5Wnu88nE as example*/
	s := Vertex{
		id:   "S",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    10,
	}
	a := Vertex{
		id:   "A",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    9,
	}
	b := Vertex{
		id:   "B",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    7,
	}
	c := Vertex{
		id:   "C",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    8,
	}
	d := Vertex{
		id:   "D",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    8,
	}
	e := Vertex{
		id:   "E",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    0,
	}
	f := Vertex{
		id:   "F",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    6,
	}
	g := Vertex{
		id:   "G",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    3,
	}
	h := Vertex{
		id:   "H",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    6,
	}
	i := Vertex{
		id:   "I",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    4,
	}
	j := Vertex{
		id:   "J",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    4,
	}
	k := Vertex{
		id:   "K",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    3,
	}
	l := Vertex{
		id:   "L",
		cost: 999999,
		adj:  make(map[interface{}]int),
		h:    6,
	}

	graph := NewGraph()
	graph.AddEdge(s, a, 7)
	graph.AddEdge(s, b, 2)
	graph.AddEdge(s, c, 3)
	graph.AddEdge(a, b, 3)
	graph.AddEdge(a, d, 4)
	graph.AddEdge(b, d, 4)
	graph.AddEdge(b, h, 1)
	graph.AddEdge(d, f, 5)
	graph.AddEdge(h, f, 3)
	graph.AddEdge(h, g, 2)
	graph.AddEdge(g, e, 2)
	graph.AddEdge(c, l, 2)
	graph.AddEdge(l, i, 4)
	graph.AddEdge(l, j, 4)
	graph.AddEdge(i, j, 6)
	graph.AddEdge(i, k, 4)
	graph.AddEdge(j, k, 4)
	graph.AddEdge(k, e, 5)

	cases := []struct {
		v, w  string
		want  int
		wpath []interface{}
	}{
		{"S", "E", 7, []interface{}{"S", "B", "H", "G", "E"}},
		/*		{1, 3, 6, []interface{}{1, 2, 3}},
				{2, 4, 7, []interface{}{2, 3, 4}},
				{1, 2, 5, []interface{}{1, 2}},
				{2, 3, 1, []interface{}{2, 3}},
				{1, 6, 999999999, nil},*/
	}
	for _, c := range cases {
		graph.ClearVisited()
		path, got := graph.AStar(c.v, c.w)
		if got != c.want {
			t.Errorf("Error: got %v cost, wanted %v", got, c.want)
		}
		if reflect.DeepEqual(path, c.wpath) == false {
			t.Errorf("Error: got %v path, wanted %v", path, c.wpath)
		}
	}
}

func TestDFSRecur(t *testing.T) {
	s := Vertex{
		id:   "S",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	a := Vertex{
		id:   "A",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	b := Vertex{
		id:   "B",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	c := Vertex{
		id:   "C",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	d := Vertex{
		id:   "D",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	e := Vertex{
		id:   "E",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	f := Vertex{
		id:   "F",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	g := Vertex{
		id:   "G",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	h := Vertex{
		id:   "H",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	i := Vertex{
		id:   "I",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	j := Vertex{
		id:   "J",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	k := Vertex{
		id:   "K",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	l := Vertex{
		id:   "L",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}

	graph := NewGraph()
	graph.AddEdge(s, a, 0)
	graph.AddEdge(s, b, 0)
	graph.AddEdge(s, c, 0)
	graph.AddEdge(a, b, 0)
	graph.AddEdge(a, d, 0)
	graph.AddEdge(b, d, 0)
	graph.AddEdge(b, h, 0)
	graph.AddEdge(d, f, 0)
	graph.AddEdge(h, f, 0)
	graph.AddEdge(h, g, 0)
	graph.AddEdge(g, e, 0)
	graph.AddEdge(c, l, 0)
	graph.AddEdge(l, i, 0)
	graph.AddEdge(l, j, 0)
	graph.AddEdge(i, j, 0)
	graph.AddEdge(i, k, 0)
	graph.AddEdge(j, k, 0)
	graph.AddEdge(k, e, 0)

	cases := []struct {
		v    string
		want bool
	}{
		{"S", true},
		//{"S", "X", false},
		/*		{1, 3, 6, []interface{}{1, 2, 3}},
				{2, 4, 7, []interface{}{2, 3, 4}},
				{1, 2, 5, []interface{}{1, 2}},
				{2, 3, 1, []interface{}{2, 3}},
				{1, 6, 999999999, nil},*/
	}
	/*for this test, just want to make sure the path is dfs in console*/
	for _, c := range cases {
		graph.ClearVisited()
		got := graph.DFS(c.v)
		if got != c.want {
			t.Errorf("Error: got %v vertex, wanted %v", got, c.want)
		}
	}
}

func TestDFSIter(t *testing.T) {
	s := Vertex{
		id:   "S",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	a := Vertex{
		id:   "A",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	b := Vertex{
		id:   "B",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	c := Vertex{
		id:   "C",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	d := Vertex{
		id:   "D",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	e := Vertex{
		id:   "E",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	f := Vertex{
		id:   "F",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	g := Vertex{
		id:   "G",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	h := Vertex{
		id:   "H",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	i := Vertex{
		id:   "I",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	j := Vertex{
		id:   "J",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	k := Vertex{
		id:   "K",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}
	l := Vertex{
		id:   "L",
		cost: 999999,
		adj:  make(map[interface{}]int),
	}

	graph := NewGraph()
	graph.AddEdge(s, a, 0)
	graph.AddEdge(s, b, 0)
	graph.AddEdge(s, c, 0)
	graph.AddEdge(a, b, 0)
	graph.AddEdge(a, d, 0)
	graph.AddEdge(b, d, 0)
	graph.AddEdge(b, h, 0)
	graph.AddEdge(d, f, 0)
	graph.AddEdge(h, f, 0)
	graph.AddEdge(h, g, 0)
	graph.AddEdge(g, e, 0)
	graph.AddEdge(c, l, 0)
	graph.AddEdge(l, i, 0)
	graph.AddEdge(l, j, 0)
	graph.AddEdge(i, j, 0)
	graph.AddEdge(i, k, 0)
	graph.AddEdge(j, k, 0)
	graph.AddEdge(k, e, 0)

	cases := []struct {
		v, w string
		want bool
	}{
		{"S", "E", true},
		{"S", "X", false},
	}
	/*for this test, just want to make sure the path is dfs in console*/
	for i, c := range cases {
		fmt.Println("\nCase ", i)
		graph.ClearVisited()
		vertex, got := graph.DFSIter(c.v, c.w)
		if got != c.want {
			t.Errorf("Error: got %v vertex %v bool, wanted %v", vertex, got, c.want)
		}
	}
}
