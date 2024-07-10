package model

type IdReq struct {
	Id int64 `json:"id" form:"id" binding:"required,gt=0" remark:"id"`
}

type IdListReq struct {
	IdList []int64 `json:"idList" form:"idList" binding:"required,gt=0" remark:"id列表"`
}
