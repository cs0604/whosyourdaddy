package kth_smallest_element_in_a_sorted_matrix

import "container/heap"

// An Item is something we manage in a priority queue.
type Item struct {
	x, y     int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func kthSmallest(matrix [][]int, k int) int {

	N := len(matrix)

	minheap := make(PriorityQueue, N)

	//first row
	for i := 0; i < N; i++ {
		item := &Item{
			x:        0,
			y:        i,
			priority: matrix[0][i],
			index:    i,
		}
		minheap[i] = item
	}

	heap.Init(&minheap)

	for i := 0; i < k-1; i++ {
		item := heap.Pop(&minheap).(*Item)
		if item.x == N-1 {
			continue
		}
		heap.Push(&minheap, &Item{
			x:        item.x + 1,
			y:        item.y,
			priority: matrix[item.x+1][item.y],
		})
	}

	return heap.Pop(&minheap).(*Item).priority
}
