package model

import (
	"time"
)

type File struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"pretty_type"`
	Size    float64 `json:"size"`
	Created int64   `json:"created"`
}

func (f File) CreatedDateWithoutTime() time.Time {
	date := time.Unix(f.Created, 0)
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	return date
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

type AuthResponse struct {
	Ok    bool   `json:"ok"`
	User  string `json:"user_id"`
	Error string `json:"error"`
}
