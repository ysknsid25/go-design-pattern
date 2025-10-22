package main

import (
	"fmt"
	"strings"
)

type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("一般的な挨拶")
	d.builder.MakeItems([]string{"Hello", "Hi", "How are you?"})

	d.builder.MakeString("時間帯に応じた挨拶")
	d.builder.MakeItems([]string{"Good morning", "Good afternoon", "Good evening"})

	d.builder.Close()
}

type TextBuilder struct {
	sb *strings.Builder
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{sb: &strings.Builder{}}
}

func (b *TextBuilder) MakeTitle(title string) {
	b.sb.WriteString("=======\n")
	b.sb.WriteString("『")
	b.sb.WriteString(title)
	b.sb.WriteString("』\n")
}

func (b *TextBuilder) MakeString(str string) {
	b.sb.WriteString("***")
	b.sb.WriteString(str)
	b.sb.WriteString("\n\n")
}

func (b *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		b.sb.WriteString("  ・")
		b.sb.WriteString(item)
		b.sb.WriteString("\n")
	}
	b.sb.WriteString("\n")
}

func (b *TextBuilder) Close() {
	b.sb.WriteString("=======\n")
}

func (b *TextBuilder) TextResult() string {
	return b.sb.String()
}

func ExecBuilder() {
	builder := NewTextBuilder()
	director := Director{builder}
	director.Construct()
	result := builder.TextResult()
	fmt.Println(result)
}
