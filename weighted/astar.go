package weighted

import "fmt"

func H(v Vertex) int {
	//for simplicity's sake, hard code heuristic per vertex
	//could be changed as needed
	return v.h
}

func (g *Graph) AStar(start, dest interface{}) ([]interface{}, int) {
	/*
		Note: all vertex distances should be set to "infitity"
		when vertices are initialized. see graph_test.go
		for graph_test, we are hardcoding the heuristic val.
		Make sure visited slice is cleared for subsequent runs.
	*/
	src := start
	g.visited[src] = true
	v := g.vertices[src]
	pq := make(PriorityQueue, 0)
	fmt.Println("At vertex: ", src)
	//init heap with adjacent vertices to src
	for id, weight := range v.adj {
		x := g.vertices[id]
		//normally we'd add distance, but from src to neighbor,
		//distance is just weight
		x.cost = weight + H(x)
		fmt.Printf("New cost for vertex %v is %v from %v\n", id, x.cost, src)
		x.from = src
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

			//when considering neighbors, we only care about the
			//neighbor's heursitic value, so we subtract
			//the current vertex's heuristic value.
			distance := weight + v.cost - H(v)
			//fmt.Printf("Distance for %v is now %v\n", w, distance)
			tentativeGScore := distance + H(neighbor)
			if tentativeGScore < neighbor.cost {
				neighbor.cost = tentativeGScore
				fmt.Printf("New cost for vertex %v is %v from %v\n", w, neighbor.cost, src)
				neighbor.from = src
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
