package main

import "fmt"

type Product interface {
	Use()
}

type ProductManager interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

type Factory struct {
	creator ProductManager
}

func (factory *Factory) Create(owner string) Product {
	product := factory.creator.CreateProduct(owner)
	factory.creator.RegisterProduct(product)
	return product
}

type IDCard struct {
	owner string
}

func (id *IDCard) Use() {
	fmt.Println(id.owner + "を使います。")
}

func (id *IDCard) GetOwner() string {
	return id.owner
}

func (id *IDCard) ToString() {
	fmt.Println("IDCard: " + id.owner)
}

func NewIDCard(owner string) *IDCard {
	return &IDCard{owner: owner}
}

type IDCardFactory struct{}

func (f *IDCardFactory) CreateProduct(owner string) Product {
	return NewIDCard(owner)
}

func (f *IDCardFactory) RegisterProduct(product Product) {
	fmt.Println("Productを登録しました。")
}

func execFactoryMethod() {
	factory := &Factory{creator: &IDCardFactory{}}
	card1 := factory.Create("山田太郎")
	card2 := factory.Create("鈴木花子")
	card3 := factory.Create("佐藤次郎")

	card1.Use()
	card2.Use()
	card3.Use()
}
