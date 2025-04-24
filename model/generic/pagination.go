package generic

import "github.com/chaihaobo/boice-blog-api/constant"

type Pagination struct {
	Page int `json:"page" form:"page" `
	Size int `json:"size" form:"size"`
}

func (p *Pagination) SetupDefault() {
	if p.Page <= 0 {
		p.Page = constant.DefaultPage
	}
	if p.Size <= 0 {
		p.Size = constant.DefaultPageSize
	}
	if p.Page > constant.MaxPageSize {
		p.Page = constant.MaxPageSize
	}

}
