package state

import (
	"github.com/t10471/go-examples/state-examples/functional/action"
	"github.com/t10471/go-examples/state-examples/functional/actiontype"
)

type VendingMachine struct {
	CurrentState State
	ItemCount    int
	ItemPrice    int
	Errs         []error
}

type State interface {
	Next(*VendingMachine, actiontype.Action) *VendingMachine
}

func NewVendingMachine(ItemCount, ItemPrice int) *VendingMachine {
	var s State
	s = &action.HasItemState{}
	if ItemCount == 0 {
		s = &action.NoItemState{}
	}
	return &VendingMachine{
		CurrentState: s,
		ItemCount:    ItemCount,
		ItemPrice:    ItemPrice,
		Errs:         nil,
	}
}

func (m *VendingMachine) Next(action actiontype.Action) *VendingMachine {
	return m.CurrentState.Next(m, action)
}

func (m *VendingMachine) GetError() error {
	return m.Errs[0]
}

func (m *VendingMachine) GetErrors() []error {
	return m.Errs
}
