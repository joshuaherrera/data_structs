package quicksort

func Sort(arr []int, lo int, hi int) {
	if lo < hi {
		p := partition(arr, lo, hi)
		Sort(arr, lo, p)
		Sort(arr, p+1, hi)
	}
}

//use Hoare's scheme
func partition(arr []int, lo int, hi int) int {
	pivot := arr[lo+(hi-lo)/2]
	for {
		for arr[lo] < pivot {
			lo++
		}
		for arr[hi] > pivot {
			hi--
		}

		if lo >= hi {
			return hi
		}
		arr[lo], arr[hi] = arr[hi], arr[lo]

		lo++
		hi--
	}
}
