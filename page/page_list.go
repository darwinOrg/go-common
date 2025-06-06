package page

import "encoding/json"

type PageList[T any] struct {
	PageNo     int  `json:"pageNo" binding:"required" remark:"页码"`
	PageSize   int  `json:"pageSize" binding:"required" remark:"每页记录数"`
	TotalCount int  `json:"totalCount" remark:"总记录数"`
	TotalPages int  `json:"totalPages" remark:"总页数"`
	List       []*T `json:"list" remark:"列表记录"`
	Extra      any  `json:"extra,omitempty" remark:"额外数据"`
}

func (p *PageList[T]) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func BuildPageList[T any](pp *PageParam, totalCount int, list []*T) *PageList[T] {
	return &PageList[T]{
		PageNo:     pp.PageNo,
		PageSize:   pp.PageSize,
		TotalCount: totalCount,
		TotalPages: calcTotalPages(totalCount, pp.PageSize),
		List:       list,
	}
}

func ListOf[T any](pageNo int, pageSize int, totalCount int, list []*T) *PageList[T] {
	return &PageList[T]{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: calcTotalPages(totalCount, pageSize),
		List:       list,
	}
}

func EmptyPageList[T any](pageNo int, pageSize int) *PageList[T] {
	return ListOf[T](pageNo, pageSize, 0, []*T{})
}

func calcTotalPages(totalCount int, pageSize int) int {
	if totalCount == 0 {
		return 0
	}

	totalPages := totalCount / pageSize

	if totalCount%pageSize > 0 {
		totalPages++
	}

	return totalPages
}
