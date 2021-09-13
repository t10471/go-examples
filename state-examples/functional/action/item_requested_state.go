package action

import (
	"fmt"

	"github.com/t10471/go-examples/state-examples/functional/actiontype"
	"github.com/t10471/go-examples/state-examples/functional/state"
)

type ItemRequestedState struct{}

func (i *ItemRequestedState) Next(st *state.VendingMachine, action actiontype.Action) *state.VendingMachine {
	switch action.ActionType {
	case actiontype.ActionInsertMoney:
		if action.InsertMoney.Money < st.ItemPrice {
			st.Errs = append(st.Errs, fmt.Errorf("inserted money is less. Please insert %d, in ItemRequestedState", st.ItemPrice))
			return st
		}
		fmt.Println("Money entered is ok")
		return &state.VendingMachine{
			CurrentState: &HasMoneyState{},
			ItemCount:    st.ItemCount,
			ItemPrice:    st.ItemPrice,
			Errs:         st.Errs,
		}
	case actiontype.ActionRequestItem, actiontype.ActionAddItem:
		st.Errs = append(st.Errs, fmt.Errorf("item Dispense in progress, in ItemRequestedState"))
		return st
	case actiontype.ActionDispenseItem:
		st.Errs = append(st.Errs, fmt.Errorf("please insert money first"))
		return st
	}
	return nil
}
