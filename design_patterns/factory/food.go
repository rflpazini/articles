package factory

import "fmt"

type FoodType int

const (
	PizzaType FoodType = iota
	SaladType
)

type Food interface {
	Prepare()
}

type Pizza struct{}

func (p Pizza) Prepare() {
	fmt.Println("Preparing a Pizza...")
}

type Salad struct{}

func (s Salad) Prepare() {
	fmt.Println("Preparing a Salad...")
}
