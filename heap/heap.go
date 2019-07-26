package heap

import (
	"fmt"
	"math"
)

//as a bi max heap
type BinaryHeap struct {
	arr  []int
	size int
}

/*
could have array above of nodes to store
values different than an int
type Node struct {
	key int
	value int
}
*/
func (heap *BinaryHeap) MaxHeapify(idx int) {
	/*
	  restores heap property, starting from a parent nod,
	  bubbling downward
	*/
	//should heapify in place
	fmt.Println("In MaxHeapify")
	left := (2 * idx) + 1
	right := (2 * idx) + 2
	largest := idx
	fmt.Println("left: ", left, "right: ", right, "largest: ", largest)
	if left <= heap.size-1 && heap.arr[left] > heap.arr[largest] {
		largest = left
	}
	if right <= heap.size-1 && heap.arr[right] > heap.arr[largest] {
		largest = right
	}
	if largest != idx {
		heap.arr[idx], heap.arr[largest] = heap.arr[largest], heap.arr[idx]
		heap.MaxHeapify(largest)
	}
}

func (heap *BinaryHeap) BubbleUp(cIdx int) {
	/*
	  Restore max heap property from a child node, bubbling
	  up until parent value > child value
	*/
	var pIdx int
	for {
		//calculate parent idx
		if (cIdx-2)%2 == 0 {
			pIdx = (cIdx - 2) / 2
		} else {
			pIdx = (cIdx - 1) / 2
		}
		if pIdx < heap.size && pIdx >= 0 {
			if heap.arr[pIdx] >= heap.arr[cIdx] {
				//done
				return
			} else {
				fmt.Println("Swapping parent and child values")
				//swap
				heap.arr[pIdx], heap.arr[cIdx] = heap.arr[cIdx], heap.arr[pIdx]
				cIdx = pIdx
			}
		} else {
			return
		}
	}
}

func (heap *BinaryHeap) Insert(val int) {
	/*inserts a new value into the max heap*/
	fmt.Println("In Inserion")
	if heap.size < 0 {
		return
	}
	heap.arr = append(heap.arr, val)
	heap.size++
	//fix heap property
	childIndex := heap.size - 1
	heap.BubbleUp(childIndex)

}

func (heap *BinaryHeap) ExtractMax() int {
	/*returns the max value and restores heap property
	  with max heapify
	*/
	max := heap.arr[0]
	//replace root with last element
	heap.arr[0] = heap.arr[heap.size-1]
	fmt.Println(len(heap.arr))
	heap.arr = heap.arr[:heap.size-1]
	fmt.Println(len(heap.arr))
	heap.size--
	fmt.Println("heap size is ", heap.size)
	heap.MaxHeapify(0)
	return max

}

func (heap *BinaryHeap) Clear() {
	heap.size = 0
	heap.arr = nil
}

func (heap *BinaryHeap) BuildMaxHeap(arr []int) {
	/*build a max heap given a slice*/
	heap.size = len(arr)
	heap.arr = arr
	for i := int(math.Floor(float64(len(arr) / 2))); i >= 0; i-- {
		heap.MaxHeapify(i)
	}
}

func (heap *BinaryHeap) PrintArr() {
	fmt.Println(heap.arr)
}
