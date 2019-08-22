package weighted

type Heap []Vertex

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) IsEmpty() bool      { return len(h) == 0 }

func (h *Heap) Push(v Vertex) {
	var changed bool
	old := *h
	updated := *h

	for i, w := range old {
		if v.id == w.id {
			//erase old v position because slice was changed
			if changed {
				if i+1 < len(updated) {
					updated = append(updated[:i], updated[i+1:]...)
				} else {
					updated = updated[:i]
				}
			} else if v.dist < w.dist {
				updated[i] = v
			}
			changed = true
		} else if v.dist < w.dist {
			//insert v before w in the slice
			changed = true
			updated = append(old[:i], v)
			updated = append(updated, w)
			updated = append(updated, old[i+1:]...)
		}
	}
	if !changed {
		updated = append(old, v)
	}
	*h = updated
}

func (h *Heap) Pop() (v Vertex) {
	old := *h
	v = old[0]
	*h = old[1:]
	return v
}

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
	heap := make(Heap, len(v.adj))
	//init heap with adjacent vertices to src
	for id, weight := range v.adj {
		v := g.vertices[id]
		v.dist = weight
		g.vertices[id] = v
		heap.Push(v)
		//fmt.Println(heap)
	}

	for src != dest {
		if heap.IsEmpty() {
			//no route
			return 999999999
		}
		v = heap.Pop()
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
			heap.Push(neighbor)
		}
	}
	v = g.vertices[dest]
	return v.dist
}
