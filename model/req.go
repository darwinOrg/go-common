package model

type IdReq struct {
	Id int64 `json:"id" binding:"required,gt=0"`
}

type IdListReq struct {
	IdList []int64 `json:"idList" binding:"required,gt=0"`
}

type StringReq struct {
	Value string `json:"value" binding:"required"`
}

type StringListReq struct {
	Values string `json:"values" binding:"required"`
}
