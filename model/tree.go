package model

type Tree struct {
	Id       int64   `json:"id,omitempty" remark:"id"`
	Name     string  `json:"name,omitempty" remark:"名称"`
	Children []*Tree `json:"children,omitempty" remark:"子节点列表"`
}
