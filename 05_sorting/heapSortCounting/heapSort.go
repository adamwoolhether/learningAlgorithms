package main

import "fmt"

type Heap struct {
	array       []int
	number      int
	swaps       int
	comparisons int
}

func NewHeap(a []int) *Heap {
	newHeap := &Heap{
		array:       a,
		number:      len(a),
		swaps:       0,
		comparisons: 0,
	}

	for k := len(a) / 2; k > 0; k-- {
		newHeap.sink(k)
	}

	return newHeap
}

func (h *Heap) sink(parent int) {
	for 2*parent <= h.number {
		child := 2 * parent
		if child < h.number && h.less(child, child+1) {
			child++
		}
		if !h.less(parent, child) {
			break
		}
		h.swap(child, parent)
		parent = child
	}
}

func (h *Heap) sort() {
	for h.number > 1 {
		h.swap(1, h.number)
		h.number--
		h.sink(1)
	}
}

func (h *Heap) less(i, j int) bool {
	h.comparisons++
	return h.array[i-1] < h.array[j-1]
}

func (h *Heap) swap(i, j int) {
	h.swaps++
	h.array[i-1], h.array[j-1] = h.array[j-1], h.array[i-1]
}

func main() {
	newHeap := NewHeap([]int{14, 13, 12, 5, 10, 6, 14, 12, 9, 1, 11, 8, 15, 9, 7, 4, 8, 2})
	fmt.Println(newHeap.array)
	fmt.Printf("swaps: %d, comparisons: %d\n", newHeap.swaps, newHeap.comparisons)
	newHeap.sort()
	fmt.Println(newHeap.array)
	fmt.Printf("swaps: %d, comparisons: %d\n", newHeap.swaps, newHeap.comparisons)

}
