package action

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/functional/actiontype"
	"github.com/t10471/go-examples/state-examples/functional/state"
)

type HasMoneyState struct{}

func (i *HasMoneyState) Next(st *state.VendingMachine, action actiontype.Action) *state.VendingMachine {
	switch action.ActionType {
	case actiontype.ActionDispenseItem:
		fmt.Println("Dispensing Item")
		st.ItemCount = st.ItemCount - 1
		var s state.State
		if st.ItemCount == 0 {
			s = &NoItemState{}
		} else {
			s = &HasItemState{}
		}
		return &state.VendingMachine{
			CurrentState: s,
			ItemCount:    st.ItemCount,
			ItemPrice:    st.ItemPrice,
			Errs:         st.Errs,
		}
	case actiontype.ActionRequestItem, actiontype.ActionAddItem:
		st.Errs = append(st.Errs, fmt.Errorf("item dispense in progress, in HasMoneyState"))
		return st
	case actiontype.ActionInsertMoney:
		st.Errs = append(st.Errs, fmt.Errorf("item out of stock, in HasMoneyState"))
		return st
	}
	return nil
}
