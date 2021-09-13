package state

import (
	"fmt"
)

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}
func (i *hasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}
