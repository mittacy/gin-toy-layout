package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/app/dp"
	"github.com/mittacy/gin-toy-layout/app/service"
	"github.com/mittacy/gin-toy-layout/app/validator/userVdr"
	"github.com/mittacy/gin-toy/core/response"
	"github.com/mittacy/gin-toy/core/singleton"
)

var User userApi

type userApi struct {
	dp dp.UserDP
}

func init() {
	singleton.Register(func() {
		User = userApi{
			dp: dp.NewUserDP(),
		}
	})
}

func (ctl *userApi) Get(c *gin.Context) {
	req := userVdr.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}

	example, err := service.User.GetById(c, req.Id)
	if err != nil {
		response.FailCheckBizErr(c, "查询记录错误", err)
		return
	}

	ctl.dp.Get(c, example)
}
