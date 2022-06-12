package exampleVdr

type GetReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}
