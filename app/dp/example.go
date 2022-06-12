package dp

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy/core/response"
)

type ExampleDP struct{}

func NewExampleDP() ExampleDP {
	return ExampleDP{}
}

func (ctl *ExampleDP) Get(c *gin.Context, data interface{}) {
	response.Success(c, data)
}
