package actiontype

type ActionType int8

const (
	ActionRequestItem  ActionType = 1
	ActionAddItem      ActionType = 2
	ActionInsertMoney  ActionType = 3
	ActionDispenseItem ActionType = 4
)

type RequestItem struct{}

type AddItem struct {
	Count int
}

type InsertMoney struct {
	Money int
}

type DispenseItem struct{}

type Action struct {
	ActionType   ActionType
	RequestItem  RequestItem
	AddItem      AddItem
	InsertMoney  InsertMoney
	DispenseItem DispenseItem
}
