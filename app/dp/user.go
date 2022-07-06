package dp

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy-layout/app/service/smodel"
	"github.com/mittacy/gin-toy-layout/app/validator/userVdr"
	"github.com/mittacy/gin-toy-layout/utils/timeUtil"
)

type UserDP struct{}

func NewUserDP() UserDP {
	return UserDP{}
}

func (ctl *UserDP) Get(c *gin.Context, data *smodel.GetById) map[string]interface{} {
	user := userVdr.GetReplyUser{
		Id:        data.User.Id,
		Name:      data.User.Name,
		Age:       data.User.Age,
		CreatedAt: timeUtil.Format(data.User.CreatedAt),
		UpdatedAt: timeUtil.Format(data.User.UpdatedAt),
	}

	return map[string]interface{}{
		"user": user,
	}
}
