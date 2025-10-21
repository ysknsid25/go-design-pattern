package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	name string
}

var (
	instance *Singleton
	once     sync.Once
)

func Instance() *Singleton {
	once.Do(func() {
		fmt.Println("インスタンスを生成しました。")
		instance = &Singleton{name: "the-one-and-only"}
	})
	return instance
}

func (s *Singleton) SetName(name string) {
	s.name = name
}

func (s *Singleton) Name() string {
	return s.name
}

func ExecSingleton() {
	fmt.Println("--- Singleton Pattern ---")
	s1 := Instance()
	s1.SetName("singleton-instance-1")

	fmt.Printf("s1.Name() = %s\n", s1.Name())

	s2 := Instance()
	fmt.Printf("s2.Name() = %s\n", s2.Name())

	if s1 == s2 {
		fmt.Println("s1とs2は同じインスタンスです。")
	}
	fmt.Println("-------------------------")
}
