package weighted

import "fmt"

/*func containsVertex(arr []interface{}, v interface{}) bool {
	for _, w := range arr {
		if v == w {
			return true
		}
	}
	return false
}*/

func (g *Graph) DFS(curr interface{}) bool {
	fmt.Printf("%v ", curr)
	//visited = append(visited, curr)
	g.visited[curr] = true
	v := g.vertices[curr]
	/*	if curr == goal {
		fmt.Println("\nVertex found.")
		return true
	}*/
	for k := range v.adj {
		//check if the node has been visited
		if _, ok := g.visited[k]; ok == false {
			g.DFS(k, goal)
		}
	}
	return true
}
