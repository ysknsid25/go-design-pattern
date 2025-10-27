package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Memento struct {
	money  int
	fruits []string
}

func (m *Memento) GetMoney() int {
	return m.money
}

func (m *Memento) GetFruits() []string {
	fruits := make([]string, len(m.fruits))
	copy(fruits, m.fruits)
	return fruits
}

type Gamer struct {
	money  int
	fruits []string
	rand   *rand.Rand
}

func NewGamer(money int) *Gamer {
	return &Gamer{
		money:  money,
		fruits: make([]string, 0),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Gamer) GetMoney() int {
	return g.money
}

func (g *Gamer) Bet() {
	dice := g.rand.Intn(6) + 1 // 1-6のサイコロ

	switch dice {
	case 1:
		g.money += 100
		fmt.Println("所持金が増えました。")
	case 2:
		g.money += 50
		fmt.Println("所持金が少し増えました。")
	case 6:
		fruit := g.getFruit()
		fmt.Printf("フルーツ（%s）をもらいました。\n", fruit)
		g.fruits = append(g.fruits, fruit)
	default:
		fmt.Println("何も起こりませんでした。")
	}
}

func (g *Gamer) CreateMemento() *Memento {
	m := &Memento{
		money: g.money,
	}
	for _, fruit := range g.fruits {
		if fruit == "りんご" {
			m.fruits = append(m.fruits, fruit)
		}
	}
	return m
}

func (g *Gamer) RestoreMemento(memento *Memento) {
	g.money = memento.money
	g.fruits = memento.GetFruits()
}

func (g *Gamer) String() string {
	return fmt.Sprintf("[money = %d, fruits = %v]", g.money, g.fruits)
}

func (g *Gamer) getFruit() string {
	prefix := ""
	if g.rand.Intn(2) == 0 {
		prefix = "おいしい"
	}

	fruitTypes := []string{"りんご", "ぶどう", "ばなな", "みかん"}
	fruit := fruitTypes[g.rand.Intn(len(fruitTypes))]

	return prefix + fruit
}

type Caretaker struct {
	mementos []*Memento
}

func NewCaretaker() *Caretaker {
	return &Caretaker{
		mementos: make([]*Memento, 0),
	}
}

func (c *Caretaker) AddMemento(memento *Memento) {
	c.mementos = append(c.mementos, memento)
	fmt.Printf("メメントを保存しました。（保存数: %d）\n", len(c.mementos))
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index >= 0 && index < len(c.mementos) {
		return c.mementos[index]
	}
	return nil
}

func (c *Caretaker) GetLatestMemento() *Memento {
	if len(c.mementos) > 0 {
		return c.mementos[len(c.mementos)-1]
	}
	return nil
}

func (c *Caretaker) GetMementoCount() int {
	return len(c.mementos)
}

func ExecMemento() {
	fmt.Println("=== Memento Pattern Demo ===")

	gamer := NewGamer(100)
	fmt.Printf("初期状態: %s\n", gamer.String())

	caretaker := NewCaretaker()

	caretaker.AddMemento(gamer.CreateMemento())

	fmt.Println("\n--- ゲーム開始 ---")

	for i := 0; i < 100; i++ {
		fmt.Printf("\n==== %d回目 ====\n", i+1)
		fmt.Printf("現在の状態: %s\n", gamer.String())

		gamer.Bet()

		fmt.Printf("所持金は%d円になりました。\n", gamer.GetMoney())

		if gamer.GetMoney() > 100 {
			fmt.Println("（だいぶ増えたので、現在の状態を保存しておこう）")
			caretaker.AddMemento(gamer.CreateMemento())
		} else if gamer.GetMoney() < 100 {
			fmt.Println("（だいぶ減ったので、以前の状態に復帰しよう）")
			latestMemento := caretaker.GetLatestMemento()
			if latestMemento != nil {
				gamer.RestoreMemento(latestMemento)
				fmt.Printf("復帰後の状態: %s\n", gamer.String())
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("\n--- ゲーム終了 ---")
	fmt.Printf("最終状態: %s\n", gamer.String())
	fmt.Printf("保存されたメメント数: %d\n", caretaker.GetMementoCount())

	fmt.Println("\n--- 保存されたメメントの履歴 ---")
	for i := 0; i < caretaker.GetMementoCount(); i++ {
		memento := caretaker.GetMemento(i)
		if memento != nil {
			fmt.Printf("メメント %d: お金=%d, フルーツ=%v\n",
				i+1, memento.GetMoney(), memento.GetFruits())
		}
	}

	fmt.Println("\n=== Demo completed ===")
}
