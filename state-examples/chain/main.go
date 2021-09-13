package main

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/chain/state"
	"github.com/t10471/go-examples/state-examples/chain/statemachine"
)

func main() {
	{
		st := &statemachine.Machine{
			State:     &state.HasItemState{},
			ItemCount: 10,
			ItemPrice: 100,
			Errs:      nil,
		}
		machine := statemachine.NewVendingMachine(st)
		r := machine.RequestItem().InsertMoney(100).DispenseItem()
		if err := r.GetError(); err != nil {
			fmt.Println(nil)
		}
	}

	{
		st := &statemachine.Machine{
			State:     &state.HasItemState{},
			ItemCount: 0,
			ItemPrice: 100,
			Errs:      nil,
		}
		machine := statemachine.NewVendingMachine(st)
		r := machine.AddItem(2).RequestItem().InsertMoney(10).DispenseItem()
		if err := r.GetError(); err != nil {
			fmt.Println(nil)
		}
	}
}
