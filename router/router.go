package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/app/api"
	"github.com/mittacy/gin-toy/core/response"
)

const version = "v0.0.1"

func Init(r *gin.Engine) {
	r.GET("/", ping)
	r.GET("/ping", ping)
	r.GET("/health", ping)

	r.GET("/get", api.Example.Get)
}

func ping(c *gin.Context) {
	res := map[string]interface{}{
		"version": version,
	}
	response.Success(c, res)
}
