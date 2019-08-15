package radixsort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{50, 40, 30, 20, 10}, []int{10, 20, 30, 40, 50}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		//{[]int{9, 88, 5, 1, 3, 0, 11, -11, -50, -999, 777, 86, 13, 7, 21}, []int{-999, -50, -11, 0, 1, 3, 5, 7, 9, 11, 13, 21, 86, 88, 777}},
		{[]int{33, 43, 56, 12, 22, 19, 27, 37, 49, 59}, []int{12, 19, 22, 27, 33, 37, 43, 49, 56, 59}},
	}

	for _, c := range cases {
		RadixSort(c.in)
		if reflect.DeepEqual(c.in, c.want) == false {
			t.Errorf("Error: got %v, wanted %v", c.in, c.want)
		}
	}
}
