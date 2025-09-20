package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).

Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

type Human struct {
	firstName string
	lastName  string
	age       int
}

func (h *Human) GetFullName() string {
	return fmt.Sprintf("%s %s", h.firstName, h.lastName)
}

func (h *Human) GetAge() int {
	return h.age
}

type Action struct {
	Human
	steps int
}

func (a *Action) AddSteps(numSteps int) {
	if numSteps > 0 {
		a.steps += numSteps
	}
	fmt.Printf("%s walked %d steps", a.firstName, a.steps)
}

func main() {
	a := Action{
		Human: Human{
			"Kate",
			"Ivanova",
			22,
		},
		steps: 0,
	}

	fmt.Println(a.GetFullName())
	fmt.Println(a.GetAge())
	a.AddSteps(5000)
}
