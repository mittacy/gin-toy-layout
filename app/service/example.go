package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/app/data"
	"github.com/mittacy/gin-toy-layout/app/model"
	"github.com/mittacy/gin-toy/core/singleton"
	"github.com/pkg/errors"
)

// 一般情况下service应该只引用并控制自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型

// Example 服务说明注释
var Example exampleService

type exampleService struct {
	data data.Example
}

func init() {
	singleton.Register(func() {
		Example = exampleService{
			data: data.NewExample(),
		}
	})
}

func (ctl *exampleService) GetById(c *gin.Context, id int) (*model.Example, error) {
	example, err := ctl.data.Get(c, id)
	if err != nil {
		return nil, errors.WithMessage(err, "查询记录错误")
	}

	return example, nil
}
