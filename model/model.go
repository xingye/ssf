package model

type File struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"pretty_type"`
	Size    float64 `json:"size"`
	Created int64   `json:"created"`
}

type Paging struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

type ListResponse struct {
	Ok       bool   `json:"ok"`
	Error    string `json:"error,omitempty"`
	Files    []File `json:"files"`
	PageInfo Paging `json:"paging"`
}
