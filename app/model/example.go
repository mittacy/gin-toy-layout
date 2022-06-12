package model

const (
	ExampleIsDeletedNo  = 0
	ExampleIsDeletedYes = 1
)

type Example struct {
}

func (*Example) TableName() string {
	return "table_name"
}
