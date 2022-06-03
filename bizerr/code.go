package bizerr

import (
	"github.com/mittacy/gin-toy/core/bizerr"
)

var (
	// A相关错误
	ARecordNoExists = &bizerr.BizErr{Code: 100101, Msg: "A记录不存在"}

	// B相关错误
	BRecordNoExists = &bizerr.BizErr{Code: 100201, Msg: "B记录不存在"}
)
