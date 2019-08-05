package hashtable

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type Key generic.Type
type Value generic.Type

type HashTable struct {
	items []Value
}

func (ht *HashTable) Init() {
	fmt.Println("Initializing")
	a := make([]Value, 0, 10)
	ht.items = a
	fmt.Println(len(ht.items), cap(ht.items))
}
