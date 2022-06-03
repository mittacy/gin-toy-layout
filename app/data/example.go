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

func (ctl *Example) GetA() error {
	return nil
	//return errors.WithStack(bizerr.ARecordNoExists)
}

func (ctl *Example) GetB() error {
	return errors.WithStack(bizerr.BRecordNoExists)
}
