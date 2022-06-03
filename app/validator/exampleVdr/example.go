package exampleVdr

type GetReq struct {
	Id int64 `form:"id" json:"id" binding:"required"`
}
