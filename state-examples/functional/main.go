package main

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/functional/actiontype"
	"github.com/t10471/go-examples/state-examples/functional/state"
)

func main() {
	{
		machine := state.NewVendingMachine(10, 100)
		r := machine.Next(actiontype.Action{
			ActionType:  actiontype.ActionRequestItem,
			RequestItem: actiontype.RequestItem{},
		}).Next(actiontype.Action{
			ActionType:  actiontype.ActionInsertMoney,
			InsertMoney: actiontype.InsertMoney{Money: 100},
		}).Next(actiontype.Action{
			ActionType:   0,
			DispenseItem: actiontype.DispenseItem{},
		})
		if err := r.GetError(); err != nil {
			fmt.Println(nil)
		}
	}
	{
		machine := state.NewVendingMachine(0, 100)
		r := machine.Next(actiontype.Action{
			ActionType: actiontype.ActionAddItem,
			AddItem:    actiontype.AddItem{Count: 10},
		}).Next(actiontype.Action{
			ActionType:  actiontype.ActionRequestItem,
			RequestItem: actiontype.RequestItem{},
		}).Next(actiontype.Action{
			ActionType:  actiontype.ActionInsertMoney,
			InsertMoney: actiontype.InsertMoney{Money: 10},
		}).Next(actiontype.Action{
			ActionType:   0,
			DispenseItem: actiontype.DispenseItem{},
		})
		if err := r.GetError(); err != nil {
			fmt.Println(nil)
		}
		if err := r.GetError(); err != nil {
			fmt.Println(nil)
		}
	}
}
