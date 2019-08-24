package weighted

//this package is used for traversals over weighted graphs

type PriorityQueue []Vertex

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PriorityQueue) IsEmpty() bool      { return len(pq) == 0 }

func (pq *PriorityQueue) Push(v Vertex) {
	var changed bool
	old := *pq
	updated := *pq

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
	*pq = updated
}

func (pq *PriorityQueue) Pop() (v Vertex) {
	old := *pq
	v = old[0]
	*pq = old[1:]
	return v
}
