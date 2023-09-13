package page

import "encoding/json"

type PageList[T any] struct {
	PageNo     int  `json:"pageNo"`
	PageSize   int  `json:"pageSize"`
	TotalCount int  `json:"totalCount"`
	TotalPages int  `json:"totalPages"`
	List       []*T `json:"list"`
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
