package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	numsMap := map[int]int{}
	var result = make([]int, k)

	for i := range nums {
		numsMap[nums[i]]++
	}

	h := &IntHeap{}
	heap.Init(h)

	for key, value := range numsMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	for i := 0; i < k; i++ {
		result[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return result
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
