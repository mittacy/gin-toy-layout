package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/app/data"
	"github.com/mittacy/gin-toy/core/singleton"
	"github.com/pkg/errors"
)

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

func (ctl *exampleService) Get(c *gin.Context) error {
	if err := ctl.data.GetA(); err != nil {
		return errors.WithMessage(err, "查询A记录错误")
	}

	if err := ctl.data.GetB(); err != nil {
		return errors.WithMessage(err, "查询B记录错误")
	}

	return nil
}
