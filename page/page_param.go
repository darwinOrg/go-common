package page

type PageParam struct {
	PageNo   int    `json:"pageNo" form:"pageNo" binding:"required,gt=0" remark:"页码"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required,gt=0" remark:"每页记录数"`
	Sort     string `json:"sort" form:"sort" remark:"排序字段"`
	Asc      bool   `json:"asc" form:"asc" remark:"是否升序"`
}

func (pp *PageParam) GetFirstResult() int {
	return (pp.PageNo - 1) * pp.PageSize
}
