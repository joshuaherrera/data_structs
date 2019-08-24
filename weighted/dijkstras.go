package weighted

func (g *Graph) ShortestPath(src, dest interface{}) (x int) {
	/*
		Note: all vertex distances should be set to "infitity"
		when vertices are initialized. see graph_test.go
	*/
	if len(g.visited) != 0 {
		//clear map on subsequent runs for testing purposes
		g.visited = make(map[interface{}]bool)
	}
	g.visited[src] = true
	v := g.vertices[src]
	pq := make(PriorityQueue, len(v.adj))
	//init heap with adjacent vertices to src
	for id, weight := range v.adj {
		v := g.vertices[id]
		v.dist = weight
		g.vertices[id] = v
		pq.Push(v)
		//fmt.Println(pq)
	}

	for src != dest {
		if pq.IsEmpty() {
			//no route
			return 999999999
		}
		v = pq.Pop()
		src = v.id
		if g.visited[src] {
			continue
		}
		g.visited[src] = true
		for w, weight := range v.adj {
			if g.visited[w] {
				continue
			}
			neighbor := g.vertices[w]
			distance := weight + v.dist
			if distance < neighbor.dist {
				neighbor.dist = distance
				g.vertices[w] = neighbor
			}
			pq.Push(neighbor)
		}
	}
	v = g.vertices[dest]
	return v.dist
}
