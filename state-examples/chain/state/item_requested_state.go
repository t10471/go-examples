package state

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/chain/statemachine"
)

type ItemRequestedState struct{}

func (i *ItemRequestedState) RequestItem(st *statemachine.Machine) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("item already requested, in ItemRequestedState"))
	return st
}

func (i *ItemRequestedState) AddItem(st *statemachine.Machine, count int) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("item Dispense in progress, in ItemRequestedState"))
	return st
}

func (i *ItemRequestedState) InsertMoney(st *statemachine.Machine, money int) *statemachine.Machine {
	if money < st.ItemPrice {
		st.Errs = append(st.Errs, fmt.Errorf("inserted money is less. Please insert %d, in ItemRequestedState", st.ItemPrice))
		return st
	}
	fmt.Println("Money entered is ok")
	return &statemachine.Machine{
		State:     &HasMoneyState{},
		ItemCount: st.ItemCount,
		ItemPrice: st.ItemPrice,
		Errs:      st.Errs,
	}
}

func (i *ItemRequestedState) DispenseItem(st *statemachine.Machine) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("please insert money first"))
	return st
}
