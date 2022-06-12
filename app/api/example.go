package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/app/dp"
	"github.com/mittacy/gin-toy-layout/app/service"
	"github.com/mittacy/gin-toy-layout/app/validator/exampleVdr"
	"github.com/mittacy/gin-toy/core/response"
	"github.com/mittacy/gin-toy/core/singleton"
)

var Example exampleApi

type exampleApi struct {
	dp dp.ExampleDp
}

func init() {
	singleton.Register(func() {
		Example = exampleApi{
			dp: dp.NewExampleDp(),
		}
	})
}

func (ctl *exampleApi) Get(c *gin.Context) {
	req := exampleVdr.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}

	if err := service.Example.Get(c, req.Id); err != nil {
		response.FailCheckBizErr(c, "查询记录错误", err)
		return
	}

	ctl.dp.Get(c, response.NilData)
}
