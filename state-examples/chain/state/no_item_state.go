package state

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/chain/statemachine"
)

type NoItemState struct{}

func (i *NoItemState) RequestItem(st *statemachine.Machine) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("please select item first, in HasItemState"))
	return st
}

func (i *NoItemState) AddItem(st *statemachine.Machine, count int) *statemachine.Machine {
	st.ItemCount = st.ItemCount + count
	return &statemachine.Machine{
		State:     &HasItemState{},
		ItemCount: st.ItemCount + count,
		ItemPrice: st.ItemPrice,
		Errs:      st.Errs,
	}
}

func (i *NoItemState) InsertMoney(st *statemachine.Machine, money int) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("item out of stock, in NoItemState"))
	return st
}

func (i *NoItemState) DispenseItem(st *statemachine.Machine) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("item out of stock, in NoItemState"))
	return st
}
