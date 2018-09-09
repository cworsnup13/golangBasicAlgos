package main

import (
	"math"
	"fmt"
	"strings"
)

type Heap interface {
	Peek() (int, error)
	Poll() (int, error)
	Add(item int)
	HeapifyDown(int)
	HeapifyUp(int)
	Remove(item int)
	Find(int) []int
	ToString() string
}

type heap struct {
	heapContainer []int
}

func New() Heap {
	h := heap{}
	return &h
}

func (h *heap) getLeftChildIndex(parent int) int {
	return 2 * parent + 1
}

func (h *heap) getRightChildIndex(parent int) int {
	return 2 * parent + 2
}

func (h *heap) getParentIndex(child int) int {
	return int(math.Floor((float64(child) - 1) / 2))
}


func (h *heap) hasParent(child int) bool {
	return h.getParentIndex(child) >= 0
}

func (h *heap) hasLeftChild(parent int) bool {
	return h.getLeftChildIndex(parent) < int(len(h.heapContainer))
}

func (h *heap) hasRightChild(parent int) bool {
	return h.getRightChildIndex(parent) < int(len(h.heapContainer))
}

func (h *heap) leftChild(parent int) int {
	return h.heapContainer[h.getLeftChildIndex(parent)]
}

func (h *heap) rightChild(parent int) int {
	return h.heapContainer[h.getRightChildIndex(parent)]
}

func (h *heap) parent(child int) (int, error) {
	if !h.hasParent(child) {
		return 0, fmt.Errorf("no parent")
	}
	return h.heapContainer[h.getParentIndex(child)], nil
}

func (h *heap) swap(idx1, idx2 int) {
	tmp := h.heapContainer[idx2]
	h.heapContainer[idx2] = h.heapContainer[idx1]
	h.heapContainer[idx1] = tmp
}

func (h *heap) Peek() (int, error) {
	if len(h.heapContainer) == 0 {
		return 0, fmt.Errorf("no values in heap")
	}

	return h.heapContainer[0], nil
}

func (h *heap) Poll() (int, error) {
	if len(h.heapContainer) == 0 {
		return 0, fmt.Errorf("no values in heap")
	}

	if len(h.heapContainer) == 1 {
		val := h.heapContainer[0]
		h.heapContainer = h.heapContainer[1:]
		return val, nil
	}

	item := h.heapContainer[0]
	h.heapContainer[0] = h.heapContainer[len(h.heapContainer)-1]
	h.HeapifyDown(0)
	return item, nil
}

func (h *heap) Add(item int) {
	h.heapContainer = append(h.heapContainer, item)
	h.HeapifyUp(len(h.heapContainer) - 1)
}

func (h *heap) Find(item int) (foundItemIndices []int) {
	for i, v := range h.heapContainer {
		if item == v {
			foundItemIndices = append(foundItemIndices, i)
		}
	}
	return foundItemIndices
}

func (h *heap) Remove(item int) {
	numItems := h.Find(item)
	for i := 0 ; i < len(numItems); i++ {

		indexToRemove := numItems[len(numItems)-1]
		if indexToRemove == len(h.heapContainer) -1 {
			h.heapContainer = h.heapContainer[0:len(h.heapContainer)-1]
		} else {
			h.heapContainer[indexToRemove] = h.heapContainer[len(h.heapContainer)-1]
			h.heapContainer = h.heapContainer[0:len(h.heapContainer)-1]

			parentItem, _ := h.parent(indexToRemove)
			if h.hasLeftChild(indexToRemove) && (!h.hasParent(indexToRemove) || h.pairInCorrectOrder(parentItem, h.heapContainer[indexToRemove])) {
				h.HeapifyDown(indexToRemove)
			} else {
				h.HeapifyUp(indexToRemove)
			}
		}
	}
}

// TODO figure out how to do min heap or max heap
func (h *heap) pairInCorrectOrder(item1, item2 int) bool {
	return item1 <= item2
}

func (h *heap) HeapifyDown(start int) {
	// Compare the parent element to its children and swap parent with the appropriate
	// child (smallest child for MinHeap, largest child for MaxHeap).
	// Do the same for next children after swap.
	var currentIndex = start
	var nextIndex = -1
	
	for h.hasLeftChild(currentIndex){
		if h.hasRightChild(currentIndex) && h.pairInCorrectOrder(h.rightChild(currentIndex), h.leftChild(currentIndex)) {
			nextIndex = h.getRightChildIndex(currentIndex)
		} else {
			nextIndex = h.getLeftChildIndex(currentIndex)
		}
		
		if h.pairInCorrectOrder(h.heapContainer[currentIndex], h.heapContainer[nextIndex]) {
			break
		}
		
		h.swap(currentIndex, nextIndex)
		currentIndex = nextIndex
	}
}

func (h *heap) HeapifyUp(start int) {
	// Take the last element (last in array or the bottom left in a tree)
	// in the heap container and lift it up until it is in the correct
	// order with respect to its parent element.
	var currentIndex = start
	parentItem, _ := h.parent(currentIndex)
	for h.hasParent(currentIndex) && !h.pairInCorrectOrder(parentItem, h.heapContainer[currentIndex]){
		h.swap(currentIndex, h.getParentIndex(currentIndex))
		currentIndex = h.getParentIndex(currentIndex)
		parentItem, _ = h.parent(currentIndex)

	}
}

func sumSquares(cap int) int {
	var sum int
	for i := 0; i < cap; i++ {
		sum += int(math.Pow(2, float64(i)))
	}
	return sum
}

func (h *heap) ToString() string {
	heapContainer := h.heapContainer
	var rows = make([]string, int(math.Ceil(math.Log2(float64(len(heapContainer))))))
	for r := 0; r < len(rows); r++ {
		currentRow := heapContainer[sumSquares(r):sumSquares(r+1)]
		var words = make([]string, len(currentRow))
		for i, v := range currentRow {
			words[i] = fmt.Sprintf("%d", v)
		}
		rows[r] = strings.Join(words, " ")
	}
	return strings.Join(rows, "\n")

}

func main() {
	h1 := New()
	var nums = []int{1,2,3,7,17,19,36,25,100}
	for _, v := range nums {
		h1.Add(v)
	}
	fmt.Println(h1.ToString())
}