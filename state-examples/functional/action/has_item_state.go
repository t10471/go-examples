package action

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/functional/actiontype"
	"github.com/t10471/go-examples/state-examples/functional/state"
)

type HasItemState struct{}

func (i *HasItemState) Next(st *state.VendingMachine, action actiontype.Action) *state.VendingMachine {
	switch action.ActionType {
	case actiontype.ActionRequestItem:
		if st.ItemCount == 0 {
			return &state.VendingMachine{
				CurrentState: &NoItemState{},
				ItemCount:    st.ItemCount,
				ItemPrice:    st.ItemPrice,
				Errs:         st.Errs,
			}
		}
		fmt.Printf("Item requestd\n")
		return &state.VendingMachine{
			CurrentState: &ItemRequestedState{},
			ItemCount:    st.ItemCount,
			ItemPrice:    st.ItemPrice,
			Errs:         st.Errs,
		}
	case actiontype.ActionAddItem:
		return &state.VendingMachine{
			CurrentState: &HasItemState{},
			ItemCount:    st.ItemCount + action.AddItem.Count,
			ItemPrice:    st.ItemPrice,
			Errs:         st.Errs,
		}
	case actiontype.ActionInsertMoney, actiontype.ActionDispenseItem:
		st.Errs = append(st.Errs, fmt.Errorf("please select item first, in HasItemState"))
		return st
	}
	return nil
}
