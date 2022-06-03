package dp

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/response"
)

type ExampleDp struct{}

func NewExampleDp() ExampleDp {
	return ExampleDp{}
}

func (ctl *ExampleDp) Get(c *gin.Context, data interface{}) {
	response.Success(c, data)
}
