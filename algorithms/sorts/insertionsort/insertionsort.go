package insertionsort

func Sort(arr []int) {
	i := 1
	var j int
	for i < len(arr) {
		j = i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
		i++
	}
}
