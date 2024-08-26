package model

type IdReq struct {
	Id int64 `json:"id" form:"id" binding:"required" remark:"id"`
}

type IdsReq struct {
	Ids []int64 `json:"ids" form:"ids" binding:"required" remark:"id列表"`
}
