package statemachine

type State interface {
	RequestItem(machine *Machine) *Machine
	AddItem(machine *Machine, count int) *Machine
	InsertMoney(machine *Machine, money int) *Machine
	DispenseItem(machine *Machine) *Machine
}

type Machine struct {
	State
	ItemCount int
	ItemPrice int
	Errs      []error
}

type VendingMachine struct {
	currentState *Machine
}

func NewVendingMachine(machine *Machine) *VendingMachine {
	return &VendingMachine{currentState: machine}
}

func (v *VendingMachine) RequestItem() *VendingMachine {
	next := v.currentState.RequestItem(v.currentState)
	v.currentState = next
	return v
}

func (v *VendingMachine) AddItem(count int) *VendingMachine {
	next := v.currentState.AddItem(v.currentState, count)
	v.currentState = next
	return v
}

func (v *VendingMachine) InsertMoney(money int) *VendingMachine {
	next := v.currentState.InsertMoney(v.currentState, money)
	v.currentState = next
	return v
}

func (v *VendingMachine) DispenseItem() *VendingMachine {
	next := v.currentState.DispenseItem(v.currentState)
	v.currentState = next
	return v
}

func (v *VendingMachine) GetError() error {
	return v.currentState.Errs[0]
}

func (v *VendingMachine) GetErrors() []error {
	return v.currentState.Errs
}
