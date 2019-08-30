package weighted

import "fmt"

/*note: would be easier if working with pointers. Refactor*/
func (g *Graph) BFS(start, goal interface{}) (Vertex, bool) {
	v := start
	q := make([]interface{}, 0)
	g.visited[v] = true
	q = append(q, v)
	path := make([]interface{}, 0)
	for len(q) != 0 {
		v = q[0]
		q = q[1:]
		path = append(path, v)
		//fmt.Println(v, goal, q)
		if v == goal {
			//vertPtr := g.vertices[v]
			//path := g.buildPath(&vertPtr, start)
			fmt.Printf("\nPath: %v\n", path)
			//fmt.Printf("\nFrom: %v, all vertices: %v", vertPtr.from, g.vertices)
			return g.vertices[v], true
		}
		for w := range g.vertices[v].adj {
			if _, ok := g.visited[w]; !ok && !containsVertex(q, w) {
				g.visited[w] = true
				vert := g.vertices[w]
				//fmt.Prsintf("\nVertex %v comes from %v\n", w, v)
				vert.from = v
				g.vertices[w] = vert
				//fmt.Printf("\nCurrent V: %v, neighbor: %v\n", v, w)
				q = append(q, w)
			}
		}
	}
	var err Vertex
	fmt.Printf("\nPath: %v", path)
	return err, false
}
