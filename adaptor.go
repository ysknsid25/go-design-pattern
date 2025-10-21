package main

import "fmt"

type Banner struct {
	strings string
}

func (banner *Banner) showWithParen() {
	fmt.Println("(" + banner.strings + ")")
}

func (banner *Banner) showWithAster() {
	fmt.Println("*" + banner.strings + "*")
}

type Print interface {
	printWeak()
	printStrong()
}

type PrintBanner struct {
	banner Banner
}

func (printBanner *PrintBanner) printWeak() {
	printBanner.banner.showWithParen()
}

func (printBanner *PrintBanner) printStrong() {
	printBanner.banner.showWithAster()
}

func NewPrintBanner(strings string) Print {
	return &PrintBanner{
		banner: Banner{strings},
	}
}

func execAdaptor() {
	printBanner := NewPrintBanner("Hello")
	printBanner.printWeak()
	printBanner.printStrong()
}
