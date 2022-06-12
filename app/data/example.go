package data

import (
	"context"
	"github.com/mittacy/gin-toy-layout/app/model"
	"github.com/mittacy/gin-toy-layout/bizerr"
	"github.com/spf13/viper"
)

type Example struct {
	//eMysql.EGorm
	//eMongo.EMongo
	//eRedis.ERedis
	//client *thirdHttp.Client
}

func NewExample() Example {
	return Example{
		//EGorm:  eMysql.EGorm{ConfName: "localhost"},
		//EMongo: eMongo.EMongo{ConfName: "localhost", Collection: "collection_name"},
		//ERedis: eRedis.ERedis{ConfName: "localhost", DB: 0},
		//client: thirdHttp.NewClient(viper.GetString("user_server_host")),
	}
}

func (ctl *Example) Get(c context.Context, id int) (*model.Example, error) {
	if id != 1 {
		return nil, bizerr.ExampleRecordNoExists
	}

	return &model.Example{}, nil
}

/*
 * 以下为查询缓存KEY方法
 */
func (ctl *Example) cacheKeyPre() string {
	return viper.GetString("APP_NAME") + ":example"
}
