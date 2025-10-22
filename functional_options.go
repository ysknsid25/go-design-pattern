package main

import "fmt"

type Coffee struct {
	Sugar int
	Milk  bool
}

type Portion func(*Coffee)

func NewCoffee(options ...Portion) *Coffee {
	coffee := &Coffee{
		Sugar: 0,
		Milk:  false,
	}
	for _, option := range options {
		option(coffee)
	}
	return coffee
}

func Sugar(sugar int) Portion {
	return func(c *Coffee) {
		c.Sugar = sugar
	}
}

func Milk() Portion {
	return func(c *Coffee) {
		c.Milk = true
	}
}

func ExecFunctionalOptions() {
	coffee1 := NewCoffee()
	coffee2 := NewCoffee(Sugar(2))
	coffee3 := NewCoffee(Milk())
	coffee4 := NewCoffee(Sugar(1), Milk())

	fmt.Printf("Coffee1: Sugar=%d, Milk=%v\n", coffee1.Sugar, coffee1.Milk)
	fmt.Printf("Coffee2: Sugar=%d, Milk=%v\n", coffee2.Sugar, coffee2.Milk)
	fmt.Printf("Coffee3: Sugar=%d, Milk=%v\n", coffee3.Sugar, coffee3.Milk)
	fmt.Printf("Coffee4: Sugar=%d, Milk=%v\n", coffee4.Sugar, coffee4.Milk)
}
