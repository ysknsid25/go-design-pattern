package main

import "fmt"

type TemplateDisplay interface {
	open()
	print()
	close()
}

type Display struct {
	template TemplateDisplay
}

func (b *Display) Display() {
	b.template.open()
	for i := 0; i < 5; i++ {
		b.template.print()
	}
	b.template.close()
}

func NewDisplay(template TemplateDisplay) *Display {
	return &Display{template: template}
}

type CharDisplay struct {
	ch rune
}

func (d *CharDisplay) open() {
	fmt.Println("<<")
}

func (d *CharDisplay) print() {
	fmt.Println(d.ch)
}

func (d *CharDisplay) close() {
	fmt.Println(">>")
}

func NewCharDisplay(ch rune) TemplateDisplay {
	return &CharDisplay{ch}
}

type StringDisplay struct {
	str string
}

func (d *StringDisplay) open() {
	d.printLine()
}

func (d *StringDisplay) print() {
	fmt.Println("|" + d.str + "|")
}

func (d *StringDisplay) close() {
	d.printLine()
}

func (d *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < len(d.str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func NewStringDisplay(str string) TemplateDisplay {
	return &StringDisplay{str}
}

func execTemplateMethod() {
	d1 := NewDisplay(NewCharDisplay('H'))
	d2 := NewDisplay(NewStringDisplay("Hello World"))

	d1.Display()
	d2.Display()
}
