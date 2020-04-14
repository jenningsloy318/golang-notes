package main

import (
	"fmt"
)

//Array declare new Array
type Array struct {
	data []int
	size int
}

// NewArray create new array
func NewArray(capacity int) Array {

	return Array{data: make([]int, capacity), size: 0}

}

// NewDefaultArray create array with default size
func NewDefaultArray() Array {
	return Array{data: make([]int, 10), size: 0}

}

// GetArraySize get array size
func (a *Array) GetArraySize() int {

	return a.size
}

//GetCapacity  get the capacity of the array
func (a Array) GetCapacity() int {
	return cap(a.data)
}

// IsEmpty check if the size is empty
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//PrintAarray will print all elements of the array
func (a *Array) PrintAarray() {
	fmt.Printf("now the size of array is: %d, items of the array is %d\n", a.size, a.data[:a.size])
}

// AddLast will add an item to the last
func (a *Array) AddLast(element int) {

	//if a.size == cap(a.data) {
	//	fmt.Errorf("Array is full, AddLast can't add element to the Array")
	//}
	//a.data[a.size] = element

	a.Add(a.size, element)
}

// AddFirst will add an item to the first
func (a *Array) AddFirst(element int) {
	a.Add(0, element)
}

// Add will add element at the index postion
func (a *Array) Add(index int, element int) {

	if a.size == cap(a.data) {
		fmt.Errorf("Array is full, AddLast can't add element to the Array")
	}

	if index < 0 || index > a.size {
		fmt.Errorf("invalid index")
	}

	// move t.data[i] ---> t.data[i+1]
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = element
	a.size++

}

func main() {

	array := NewDefaultArray()
	array.PrintAarray()

	array.AddLast(6)
	array.PrintAarray()

	array.Add(0, 7)
	array.PrintAarray()

	array.Add(1, 8)
	array.PrintAarray()

}