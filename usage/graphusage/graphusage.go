package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/graph"
)

func main() {
	fmt.Println("Testing graph")
	g := graph.NewGraph()
	fmt.Println(g.V())
	fmt.Println(g.E())
	g.AddEdge("A", "B")
	fmt.Println(g.V())
	fmt.Println(g.E())
	g.AddEdge("C", "B")
	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println("Testing degree")
	deg, err := g.Degree("A")
	fmt.Println("A: ", deg, err)
	fmt.Println("Testing adjacency")
	list, e := g.Adj("A")
	fmt.Println("Adj of A", list, e)
	fmt.Println("\n\n", g.AdjList())

}
