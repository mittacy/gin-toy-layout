package bizerr

import (
	"github.com/mittacy/gin-toy/core/bizerr"
)

var (
	// example相关错误
	ExampleRecordNoExists = &bizerr.BizErr{Code: 100101, Msg: "example记录不存在"}
)
