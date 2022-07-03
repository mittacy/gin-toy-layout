package bizerr

import (
	"github.com/mittacy/gin-toy/core/bizerr"
)

var (
	// 用户相关错误
	UserNoExists = &bizerr.BizErr{Code: 100101, Msg: "用户不存在"}
)
