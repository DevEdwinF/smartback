package entity

import (
	"context"

	"github.com/go-playground/mold/v4/modifiers"
)

var conform = modifiers.New()

type Paginate struct {
	Page  int `query:"page" mod:"default=1"`
	Limit int `query:"limit" mod:"default=100"`
}

type Pagination struct {
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	TotalPage int         `json:"total_page"`
	TotalRows int64       `json:"total_rows"`
	Rows      interface{} `json:"rows"`
}

func (p *Paginate) SetDefault() {
	_ = conform.Struct(context.Background(), p)
}
