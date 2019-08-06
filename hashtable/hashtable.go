package hashtable

import (
	"fmt"
	"math"
)

type Key interface{}
type Value interface{}

type TableItem struct {
	key   Key
	value Value
}

type HashTable struct {
	buckets [][]TableItem
	size    int //amount of entries, used for load factor
}

const LOAD_FACTOR = .7

func (ht *HashTable) Init() {
	/*Initialize the hash table with 5 values*/
	fmt.Println("Initializing")
	/*
	** init 5 slices
	** need to have cap == len to allow space for new additions
	** will have [[][][][][]]
	** Before had [] due to using make([][]TableItem,0, 5)
	** which meant our table had 0 slots available.
	 */
	a := make([][]TableItem, 5)
	for i := range a {
		//we can have 0 slots here since we just append
		a[i] = make([]TableItem, 0, 2)
	}
	ht.buckets = a
	fmt.Println(len(ht.buckets), cap(ht.buckets))
}

func (ht *HashTable) hash(k Key) int {
	/*//use Horner's method as a hash*/
	fmt.Println("Hashing")
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = (31*h + int(key[i])) % cap(ht.buckets)
	}
	return h
}

func isPrime(i int) bool {
	/*checks if a given value is prime*/
	fmt.Println("Checking if prime")
	if i == 1 {
		return false
	}
	if i == 2 {
		return true
	}

	limit := int(math.Ceil(math.Sqrt(float64(i))))
	fmt.Println("limit is: ", limit, "i is: ", i)
	for j := 2; j <= limit; j++ {
		//fmt.Println("In loop: ", i)
		if i%j == 0 {
			//fmt.Println("not prime")
			return false
		}
	}
	return true
}

func nearestPrime(i int) int {
	/*Returns the nearest prime value to i*/
	fmt.Println("Getting nearest prime")
	if isPrime(i) == true {
		//fmt.Println("Found prime: ", i)
		return i
	}
	for {
		i++
		if isPrime(i) == true {
			//fmt.Println("Found prime: ", i)
			return i
		}
	}

}

func (ht *HashTable) resize() {
	/*resizes the slice to add more values if load factor
	  surpasses 0.7
	*/
	fmt.Println("Resizing")
	//due to Put incrementing size
	ht.size = 0
	//double the array, then find nearest prime for efficiency
	newCap := nearestPrime(len(ht.buckets) * 2)
	fmt.Println("newCap is ", newCap)

	a := make([][]TableItem, newCap)
	b := make([][]TableItem, len(ht.buckets))
	copy(b, ht.buckets)
	ht.buckets = a
	fmt.Println("\n\n", a, "\n\n", b, "\n\n", ht.buckets)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			ht.Put(b[i][j].key, b[i][j].value)
		}
	}

	b = nil
	fmt.Println("Capacity is now: ", cap(ht.buckets))
}

func (ht *HashTable) Put(k Key, v Value) {
	/*adds a value to the table using the given key*/
	fmt.Println("Putting value")
	if float64(ht.size/cap(ht.buckets)) > LOAD_FACTOR {
		ht.resize()
	}
	idx := ht.hash(k)
	fmt.Println("Hashed idx is: ", idx)
	ht.buckets[idx] = append(ht.buckets[idx], TableItem{k, v})
	ht.size++
}

func (ht *HashTable) Get(k Key) (Value, bool) {
	fmt.Println("Getting a value")
	idx := ht.hash(k)
	for i := 0; i < len(ht.buckets[idx]); i++ {
		if ht.buckets[idx][i].key == k {
			return ht.buckets[idx][i].value, true
		}
	}
	return nil, false
}

func (ht *HashTable) Remove(k Key) {
	fmt.Println("Removing a value")
	idx := ht.hash(k)
	for i := 0; i < len(ht.buckets[idx]); i++ {
		if ht.buckets[idx][i].key == k {
			ht.buckets[idx] = append(ht.buckets[idx][:i], ht.buckets[idx][i+1:]...)
		}
	}
}

func (ht *HashTable) PrintBuckets() {
	for i := 0; i < len(ht.buckets); i++ {
		fmt.Println("Bucket #: ", i)
		for j := 0; j < len(ht.buckets[i]); j++ {
			fmt.Println("\tKey: ", ht.buckets[i][j].key, "\n\tValue: ", ht.buckets[i][j].value)
		}
	}
}
