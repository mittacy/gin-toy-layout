package data

import (
	"github.com/mittacy/gin-toy-layout/bizerr"
	"github.com/pkg/errors"
)

type Example struct {
}

func NewExample() Example {
	return Example{}
}

func (ctl *Example) GetA(aId int) error {
	return nil
}

func (ctl *Example) GetB(id int) error {
	if id != 1 {
		return errors.WithStack(bizerr.BRecordNoExists)
	}
	return nil
}
