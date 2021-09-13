package state

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/chain/statemachine"
)

type HasItemState struct{}

func (i *HasItemState) RequestItem(st *statemachine.Machine) *statemachine.Machine {
	if st.ItemCount == 0 {
		return &statemachine.Machine{
			State:     &NoItemState{},
			ItemCount: st.ItemCount,
			ItemPrice: st.ItemPrice,
			Errs:      st.Errs,
		}
	}
	fmt.Printf("Item requestd\n")
	return &statemachine.Machine{
		State:     &ItemRequestedState{},
		ItemCount: st.ItemCount,
		ItemPrice: st.ItemPrice,
		Errs:      st.Errs,
	}
}

func (i *HasItemState) AddItem(st *statemachine.Machine, count int) *statemachine.Machine {
	fmt.Printf("%d items added\n", count)
	st.ItemCount = st.ItemCount + count
	return &statemachine.Machine{
		State:     &HasItemState{},
		ItemCount: st.ItemCount + count,
		ItemPrice: st.ItemPrice,
		Errs:      st.Errs,
	}
}

func (i *HasItemState) InsertMoney(st *statemachine.Machine, money int) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("please select item first, in HasItemState"))
	return st
}
func (i *HasItemState) DispenseItem(st *statemachine.Machine) *statemachine.Machine {
	st.Errs = append(st.Errs, fmt.Errorf("please select item first, in HasItemState"))
	return st
}
