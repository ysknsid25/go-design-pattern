package main

import (
	"fmt"
	"strconv"
)

type Observer interface {
	Update(generator NumberGenerator)
}

type Subject interface {
	AddObserver(observer Observer)
	DeleteObserver(observer Observer)
	NotifyObservers()
}

type NumberGenerator interface {
	Subject
	GetNumber() int
	Execute()
}

type RandomNumberGenerator struct {
	observers []Observer
	number    int
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{
		observers: make([]Observer, 0),
		number:    0,
	}
}

func (rng *RandomNumberGenerator) GetNumber() int {
	return rng.number
}

func (rng *RandomNumberGenerator) AddObserver(observer Observer) {
	rng.observers = append(rng.observers, observer)
}

func (rng *RandomNumberGenerator) DeleteObserver(observer Observer) {
	for i, obs := range rng.observers {
		if obs == observer {
			rng.observers = append(rng.observers[:i], rng.observers[i+1:]...)
			break
		}
	}
}

func (rng *RandomNumberGenerator) NotifyObservers() {
	for _, observer := range rng.observers {
		observer.Update(rng)
	}
}

func (rng *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		rng.number = int(rng.randInt(50))
		rng.NotifyObservers()
	}
}

func (rng *RandomNumberGenerator) randInt(max int) int {
	rng.number = (rng.number*1103515245 + 12345) % (1 << 31)
	if rng.number < 0 {
		rng.number = -rng.number
	}
	return rng.number % max
}

type DigitObserver struct{}

func NewDigitObserver() *DigitObserver {
	return &DigitObserver{}
}

func (do *DigitObserver) Update(generator NumberGenerator) {
	fmt.Printf("DigitObserver: %d\n", generator.GetNumber())
	try(100)
}

type GraphObserver struct{}

func NewGraphObserver() *GraphObserver {
	return &GraphObserver{}
}

func (go_ *GraphObserver) Update(generator NumberGenerator) {
	fmt.Print("GraphObserver: ")
	count := generator.GetNumber()
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	try(100) // 100ms待機
}

type IncrementalObserver struct {
	prevNumber int
}

func NewIncrementalObserver() *IncrementalObserver {
	return &IncrementalObserver{
		prevNumber: -1,
	}
}

func (io *IncrementalObserver) Update(generator NumberGenerator) {
	currentNumber := generator.GetNumber()
	if io.prevNumber != -1 {
		diff := currentNumber - io.prevNumber
		fmt.Printf("IncrementalObserver: %d (diff: %+d)\n", currentNumber, diff)
	} else {
		fmt.Printf("IncrementalObserver: %d (initial)\n", currentNumber)
	}
	io.prevNumber = currentNumber
	try(100)
}

type FrameObserver struct{}

func NewFrameObserver() *FrameObserver {
	return &FrameObserver{}
}

func (fo *FrameObserver) Update(generator NumberGenerator) {
	numberStr := strconv.Itoa(generator.GetNumber())
	width := len(numberStr) + 4

	fmt.Print("FrameObserver: +")
	for i := 0; i < width-2; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	fmt.Printf("| %s |\n", numberStr)

	fmt.Print("+")
	for i := 0; i < width-2; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	try(100)
}

func try(milliseconds int) {
	for i := 0; i < milliseconds*1000; i++ {
	}
}

func ExecObserver() {
	fmt.Println("=== Observer Pattern Demo ===")

	generator := NewRandomNumberGenerator()

	digitObserver := NewDigitObserver()
	graphObserver := NewGraphObserver()
	incrementalObserver := NewIncrementalObserver()
	frameObserver := NewFrameObserver()

	generator.AddObserver(digitObserver)
	generator.AddObserver(graphObserver)
	generator.AddObserver(incrementalObserver)
	generator.AddObserver(frameObserver)

	fmt.Println("\n--- Starting number generation ---")

	generator.Execute()

	fmt.Println("\n--- Removing GraphObserver ---")

	generator.DeleteObserver(graphObserver)
	generator.Execute()

	fmt.Println("\n=== Demo completed ===")
}
