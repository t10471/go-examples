package action

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/functional/actiontype"
	"github.com/t10471/go-examples/state-examples/functional/state"
)

type NoItemState struct{}

func (i *NoItemState) Next(st *state.VendingMachine, action actiontype.Action) *state.VendingMachine {
	switch action.ActionType {
	case actiontype.ActionRequestItem:
		st.ItemCount = st.ItemCount + action.AddItem.Count
		return &state.VendingMachine{
			CurrentState: &HasItemState{},
			ItemCount:    st.ItemCount + action.AddItem.Count,
			ItemPrice:    st.ItemPrice,
			Errs:         st.Errs,
		}
	case actiontype.ActionAddItem:
		st.Errs = append(st.Errs, fmt.Errorf("item Dispense in progress, in ItemRequestedState"))
		return st
	case actiontype.ActionInsertMoney, actiontype.ActionDispenseItem:
		st.Errs = append(st.Errs, fmt.Errorf("item out of stock, in NoItemState"))
		return st
	}
	return nil
}
