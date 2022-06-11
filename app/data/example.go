package data

import (
	"github.com/mittacy/gin-toy-layout/bizerr"
	"github.com/pkg/errors"
)

type Example struct {
	//eMysql.EGorm
	//eRedis.ERedis
	//eMongo.EMongo
}

func NewExample() Example {
	return Example{
		//EGorm:  eMysql.EGorm{ConfName: "localhost"},
		//ERedis: eRedis.ERedis{ConfName: "localhost", DB: 0},
		//EMongo: eMongo.EMongo{ConfName: "localhost", Collection: "collection_name"},
	}
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
