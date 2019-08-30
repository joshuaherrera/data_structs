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
		if v == goal {
			fmt.Printf("\nPath: %v\n", path)
			return g.vertices[v], true
		}
		for w := range g.vertices[v].adj {
			if _, ok := g.visited[w]; !ok && !containsVertex(q, w) {
				g.visited[w] = true
				vert := g.vertices[w]
				vert.from = v
				g.vertices[w] = vert
				q = append(q, w)
			}
		}
	}
	var err Vertex
	fmt.Printf("\nPath: %v", path)
	return err, false
}
