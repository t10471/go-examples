package state

import (
	"fmt"
)

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *itemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}
