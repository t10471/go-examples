package state

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/chain/statemachine"
)

type HasMoneyState struct{}

func (i *HasMoneyState) RequestItem(st *statemachine.Machine) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("item dispense in progress, in HasMoneyState"))
	return st
}

func (i *HasMoneyState) AddItem(st *statemachine.Machine, count int) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("item dispense in progress, in HasMoneyState"))
	return st
}

func (i *HasMoneyState) InsertMoney(st *statemachine.Machine, money int) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("item out of stock, in HasMoneyState"))
	return st
}

func (i *HasMoneyState) DispenseItem(st *statemachine.Machine) *statemachine.Machine {
	fmt.Println("Dispensing Item")
	st.ItemCount = st.ItemCount - 1
	var s statemachine.State
	if st.ItemCount == 0 {
		s = &NoItemState{}
	} else {
		s = &HasItemState{}
	}
	return &statemachine.Machine{
		State:     s,
		ItemCount: st.ItemCount,
		ItemPrice: st.ItemPrice,
		Errs:      st.Errs,
	}
}
