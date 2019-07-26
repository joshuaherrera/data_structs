package main

import (
	"fmt"

	"github.com/joshuaherrera/data_structs/heap"
)

func main() {
	fmt.Println("Testing BinaryHeap")
	binheap := new(heap.BinaryHeap)
	binheap.Insert(1)
	binheap.PrintArr()
	binheap.Insert(2)
	binheap.PrintArr()
	binheap.Insert(3)
	binheap.PrintArr()
	binheap.Insert(-5)
	binheap.PrintArr()
	binheap.Insert(70)
	binheap.PrintArr()
	binheap.Insert(4)
	binheap.PrintArr()
	binheap.Insert(999)
	binheap.PrintArr()
	max := binheap.ExtractMax()
	fmt.Println("Max is ", max)
	binheap.PrintArr()
	max = binheap.ExtractMax()
	fmt.Println("Max is ", max)
	binheap.PrintArr()
	max = binheap.ExtractMax()
	fmt.Println("Max is ", max)
	binheap.PrintArr()
	binheap.Clear()
	binheap.PrintArr()
	a := []int{1, 2, 3, 4, 5, 6, 7, 777}
	binheap.BuildMaxHeap(a)
	binheap.PrintArr()

}
