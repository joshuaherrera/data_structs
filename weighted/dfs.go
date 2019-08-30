package weighted

import "fmt"

func (g *Graph) DFS(curr interface{}) bool {
	fmt.Printf("%v ", curr)
	g.visited[curr] = true
	v := g.vertices[curr]
	for k := range v.adj {
		//check if the node has been visited
		if _, ok := g.visited[k]; ok == false {
			g.DFS(k)
		}
	}
	return true
}

func (g *Graph) DFSIter(start, end interface{}) (Vertex, bool) {
	/*NOTE: for subsequent runs, make sure visited slice is cleared
	  see graph_test for example
	*/
	stack := make([]interface{}, 0)
	stack = append(stack, start)
	for len(stack) != 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%v ", v)
		if v == end {
			return g.vertices[v], true
		}
		if _, ok := g.visited[v]; ok == false {
			g.visited[v] = true
			for w := range g.vertices[v].adj {
				if _, ok := g.visited[w]; !ok && containsVertex(stack, w) == false {
					stack = append(stack, w)
				}
			}
		}
	}
	var err Vertex
	return err, false
}
