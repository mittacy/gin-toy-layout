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

func (ctl *exampleService) Get(c *gin.Context, id int) error {
	if _, err := ctl.data.Get(c, id); err != nil {
		return errors.WithMessage(err, "查询记录错误")
	}

	return nil
}
