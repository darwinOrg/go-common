package model

type Tree struct {
	Id       int64   `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Children []*Tree `json:"children,omitempty"`
}
