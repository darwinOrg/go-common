package page

type PageParam struct {
	PageNo   int    `json:"pageNo" form:"pageNo" binding:"required,gt=0"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required,gt=0"`
	Sort     string `json:"sort" form:"sort"`
	Asc      bool   `json:"asc" form:"asc"`
}

func (pp *PageParam) GetFirstResult() int {
	return (pp.PageNo - 1) * pp.PageSize
}
