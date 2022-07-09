package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	elements  []int
	lastIndex int
}

var pushLock sync.Mutex

const (
	capacity = 10
)

func (stack *Stack) Push(val int) {
	pushLock.Lock()
	defer pushLock.Unlock()
	if stack.lastIndex == capacity-1 {
		return
	}
	stack.elements = append(stack.elements, val)
	stack.lastIndex += 1
	fmt.Println("Pushed new element to array:", val)
}

func (stack *Stack) Pop() {
	pushLock.Lock()
	defer pushLock.Unlock()
	if stack.lastIndex == -1 {
		return
	}
	stack.elements = stack.elements[:stack.lastIndex]
	stack.lastIndex -= 1
	fmt.Println("Popped last element from array")
}

func (stack *Stack) Display() {
	pushLock.Lock()
	defer pushLock.Unlock()
	fmt.Println("List now is:", stack.elements)
}

func main() {
	var wg sync.WaitGroup
	stack := Stack{
		elements:  make([]int, 0),
		lastIndex: -1,
	}
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 23, 34, 24}
	listLength := len(list)

	wg.Add(listLength + (listLength / 2))

	for index, val := range list {
		if index%2 == 1 {
			go func() {
				stack.Pop()
				stack.Display()
				wg.Done()
			}()
		}
		go func(v int) {
			stack.Push(v)
			stack.Display()
			wg.Done()
		}(val)
	}

	wg.Wait()

	fmt.Println("Done.")
}
