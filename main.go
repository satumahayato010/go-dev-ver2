package main

import (
	"errors"
	"fmt"
)

type Money struct {
	amount   int
	currency Currency
}

type Currency struct {
	value int
}

func NewMoney(amount int, currency Currency) (*Money, error) {
	if amount < 0 {
		return nil, errors.New("ERROR")
	}
	return &Money{amount, currency}, nil
}

/*func (m *Money) add(other int) {
	m.amount = m.amount + other
}*/

func (m *Money) add(other int) *Money {
	added := m.amount + other
	return &Money{added, m.currency}
}

func main() {
	cur := Currency{}
	mon, err := NewMoney(3, cur)
	if err != nil {
		fmt.Println("ERROR")
	}
	mon2 := mon.add(3)
	fmt.Println(mon2)
	fmt.Println(mon)
}
