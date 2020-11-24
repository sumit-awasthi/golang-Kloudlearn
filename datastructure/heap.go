//heap implementation in golang #max and min heap
package main

import "fmt"

type maxheap struct {
	heapArray []int
	size      int
	maxsize   int
}
type minheap struct {
	heapArray []int
	size      int
	maxsize   int
}

//construction of the max and the min heap
func newMaxHeap(maxsize int) *maxheap {
	maxheap := &maxheap{
		heapArray: []int{},
		size:      0,
		maxsize:   maxsize,
	}
	return maxheap
}
func newMinHeap(maxsize int) *minheap {
	minheap := &minheap{
		heapArray: []int{},
		size:      0,
		maxsize:   maxsize,
	}
	return minheap
}

//insertion in the leaf of the BST
func (m *maxheap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}
func (m *minheap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *maxheap) parent(index int) int {
	return (index - 1) / 2
}
func (m *minheap) parent(index int) int {
	return (index - 1) / 2
}

// for max heap Left Child – 2*i + 1 & Right Child – 2*i + 2

func (m *maxheap) leftchild(index int) int {
	return 2*index + 1
}

//for min heap Left Child – 2*i + 1&Right Child – 2*i + 2
func (m *minheap) leftchild(index int) int {
	return 2*index + 1
}

func (m *maxheap) rightchild(index int) int {
	return 2*index + 2
}
func (m *minheap) rightchild(index int) int {
	return 2*index + 2
}

func (m *maxheap) insert(item int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("Heal is ful")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

func (m *minheap) insert(item int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("Heal is ful")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

func (m *maxheap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}

func (m *minheap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}

func (m *maxheap) upHeapify(index int) {
	for m.heapArray[index] > m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}
func (m *minheap) upHeapify(index int) {
	for m.heapArray[index] < m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
	}
}

func (m *maxheap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := m.leftchild(current)
	rightRightIndex := m.rightchild(current)
	//If current is smallest then return
	if leftChildIndex < m.size && m.heapArray[leftChildIndex] > m.heapArray[largest] {
		largest = leftChildIndex
	}
	if rightRightIndex < m.size && m.heapArray[rightRightIndex] > m.heapArray[largest] {
		largest = rightRightIndex
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
	return
}
func (m *minheap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	smallest := current
	leftChildIndex := m.leftchild(current)
	rightRightIndex := m.rightchild(current)
	//If current is smallest then return
	if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < m.size && m.heapArray[rightRightIndex] < m.heapArray[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
	return
}

func (m *maxheap) buildMaxHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}
func (m *minheap) buildMinHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *maxheap) remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top
}

func (m *minheap) remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top
}

func main() {
	inputArray := []int{6, 15, 3, 7, 2, 18, 45, 8, 6, 4, 9}
	maxHeap := newMaxHeap(len(inputArray))
	minHeap := newMinHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		maxHeap.insert(inputArray[i])
	}
	for i := 0; i < len(inputArray); i++ {
		minHeap.insert(inputArray[i])
	}
	maxHeap.buildMaxHeap()
	fmt.Println("The Max Heap is:- ")
	for i := 0; i < len(inputArray); i++ {
		fmt.Println("The Max at level", i, "is", maxHeap.remove())
	}
	fmt.Println("")
	fmt.Println("^^^^^^^^^^^^^^^^***********^^^^^^^^^^^^^^^^^^")
	minHeap.buildMinHeap()
	fmt.Println("The MIN Heap is:-")
	for i := 0; i < len(inputArray); i++ {
		fmt.Println("The Min at level", i, "is", minHeap.remove())
	}
	fmt.Scanln()
}
