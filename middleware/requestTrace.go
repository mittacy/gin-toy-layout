package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/utils/bizUtil"
)

const (
	RequestIdKey = "requestId"
)

// RequestTrace 请求追踪，添加请求id……
// @param traceIdKey 如果追踪id key能获取到追踪id，则使用该id为追踪id，否则生成新的请求id
// @return gin.HandlerFunc
func RequestTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := bizUtil.NewRequestId()
		c.Set(RequestIdKey, requestId)
		c.Next()
	}
}
