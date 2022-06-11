package exampleVdr

type GetReq struct {
	AId int `form:"a_id" json:"a_id"`
	BId int `form:"b_id" json:"b_id" binding:"required"`
}
