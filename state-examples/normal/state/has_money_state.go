package state

import (
	"fmt"
)

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *hasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
