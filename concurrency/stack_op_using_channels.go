package main

import (
	"fmt"
	"sync"
)

const (
	PushOp = iota
	PopOp
)

type Stack1 struct {
	elements    []int
	cap         int
	lastIndex   int
	operationCh chan Operation
	popVals     chan int
}

func NewStack(size int) *Stack1 {
	newStack := &Stack1{
		elements:    make([]int, 0, size),
		cap:         size,
		lastIndex:   -1,
		operationCh: make(chan Operation),
		popVals:     make(chan int),
	}

	go newStack.ProcessOpsChannel()
	go newStack.ProcessPoppedValues()

	return newStack
}

type Operation struct {
	newVal    int
	operation int
}

func (stack *Stack1) GetLength() int {
	return stack.lastIndex + 1
}

func (stack *Stack1) IncrementIndex() {
	stack.lastIndex += 1
}

func (stack *Stack1) DecrementIndex() {
	stack.lastIndex -= 1
}

func (stack *Stack1) CloseChannels() {
	close(stack.operationCh)
	close(stack.popVals)
}

func (stack *Stack1) ProcessOpsChannel() {
	for operation := range stack.operationCh {
		switch operation.operation {
		case PushOp:
			if stack.GetLength() == stack.cap {
				fmt.Println(fmt.Sprintf("Stack overflow while trying to add: %d, skipping insert.", operation.newVal))
				continue
			}
			stack.elements = append(stack.elements, operation.newVal)
			stack.IncrementIndex()
			stack.Display()
		case PopOp:
			currentLength := stack.GetLength()
			if currentLength == 0 {
				fmt.Println(fmt.Sprintf("Stack underflow while trying to delete last element."))
				continue
			} else if currentLength == 1 {
				stack.popVals <- stack.elements[0]
				stack.elements = nil
			} else {
				popElement := stack.elements[currentLength-1]
				stack.elements = stack.elements[:currentLength-1]
				stack.popVals <- popElement
			}
			stack.DecrementIndex()
			stack.Display()
		default:
			fmt.Println(fmt.Sprintf("Received invalid operation value: %d", operation.operation))
		}

	}
}

func (stack Stack1) Display() {
	fmt.Println("Stack elements now:", stack.elements)
}

func (stack *Stack1) ProcessPoppedValues() {
	for val := range stack.popVals {
		fmt.Println(fmt.Sprintf("Popped last element: %d", val))
	}
}

func main() {
	var wg sync.WaitGroup
	newStack := NewStack(20)
	elements := []int{1, 2, 3, 4, 5, 6, 12, 23, 72, 23, 53, 55, 65, 35, 65, 83, 44, 98, 87, 24}

	wg.Add(len(elements))
	for _, val := range elements {
		if val%2 == 0 {
			go func(num int) {
				newStack.operationCh <- Operation{operation: PushOp, newVal: num}
				wg.Done()
			}(val)
		} else {
			go func(num int) {
				newStack.operationCh <- Operation{operation: PopOp}
				wg.Done()
			}(val)
		}
	}

	wg.Wait()
	newStack.CloseChannels()
}
