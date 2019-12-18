package intcode

import (
	"sync"
)

type ComputerArray struct {
	size       int
	isFeedback bool
	computers  []*Computer
	inputs     []chan int64
}

func NewArray(size int) *ComputerArray {
	arr := ComputerArray{size: size}
	for i := 0; i < size; i++ {
		arr.inputs = append(arr.inputs, make(chan int64, 100))
	}
	arr.inputs = append(arr.inputs, make(chan int64, 100))
	for i := 0; i < size; i++ {
		comp := NewComputer(arr.inputs[i], arr.inputs[i+1])
		arr.computers = append(arr.computers, comp)
	}
	return &arr
}

func NewFeedbackArray(size int) *ComputerArray {
	arr := ComputerArray{size: size, isFeedback: true}
	for i := 0; i < size; i++ {
		arr.inputs = append(arr.inputs, make(chan int64, 100))
	}
	for i := 0; i < size; i++ {
		comp := NewComputer(arr.inputs[i], arr.inputs[(i+1)%size])
		arr.computers = append(arr.computers, comp)
	}
	return &arr
}

func (a *ComputerArray) SetPhase(phases []int64) {
	for i := 0; i < a.size; i++ {
		a.inputs[i] <- phases[i]
	}
}

func (a *ComputerArray) WriteInitialInput(n int64) {
	a.inputs[0] <- n
}

func (a *ComputerArray) SetProgram(prog string) {
	for i := 0; i < a.size; i++ {
		a.computers[i].SetInput(prog)
	}
}

func (a *ComputerArray) Run() {
	var wg sync.WaitGroup

	wg.Add(a.size)
	for i := 0; i < a.size; i++ {
		go func(n int) {
			defer wg.Done()
			a.computers[n].Calculate()
		}(i)
	}

	wg.Wait()
}

func (a *ComputerArray) GetResult() int64 {
	if a.isFeedback {
		return <-a.inputs[0]
	} else {
		return <-a.inputs[a.size]
	}
}
