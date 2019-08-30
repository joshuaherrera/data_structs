package weighted

import "fmt"

func (g *Graph) ShortestPath(start, dest interface{}) ([]interface{}, int) {
	/*
		Note: all vertex distances should be set to "infitity"
		when vertices are initialized. see graph_test.go
		Also, make sure visited slice is cleared of values for
		subsequent runs.
	*/
	src := start
	g.visited[src] = true
	v := g.vertices[src]
	pq := make(PriorityQueue, 0)
	fmt.Println("At vertex: ", src)
	//init heap with adjacent vertices to src
	for id, weight := range v.adj {
		x := g.vertices[id]
		x.cost = weight
		x.from = src
		fmt.Printf("New cost for vertex %v is %v from %v\n", id, x.cost, src)
		g.vertices[id] = x
		pq.Push(x)
	}

	for src != dest {
		if pq.IsEmpty() {
			//no route
			return nil, 999999999
		}
		v = pq.Pop()
		src = v.id
		fmt.Println("At vertex: ", src)
		if g.visited[src] {
			continue
		}
		g.visited[src] = true
		for w, weight := range v.adj {
			if g.visited[w] {
				continue
			}
			neighbor := g.vertices[w]
			distance := weight + v.cost
			if distance < neighbor.cost {
				neighbor.cost = distance
				neighbor.from = src
				fmt.Printf("New cost for vertex %v is %v from %v\n", w, neighbor.cost, src)
				g.vertices[w] = neighbor

			}
			pq.Push(neighbor)
		}
	}
	v = g.vertices[dest]
	path := g.buildPath(&v, start)
	fmt.Println(path, v.cost)
	return path, v.cost
}
