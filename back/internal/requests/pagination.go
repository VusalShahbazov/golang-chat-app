package requests

type Pagination struct {
	PerPage int `json:"per_page" binding:"omitempty,min=0" form:"per_page" `
	Page int `json:"page" binding:"omitempty,min=0" form:"page"`
}

type IPagination interface {
	Limit() int
	Offset() int
}
const  PerPage  = 12

func (p *Pagination)  Limit() int {
	if p.PerPage == 0  {
		return PerPage
	}
	return  p.PerPage
}

func (p *Pagination)  Offset() int {
	if p.Page == 0  {
		return  0
	}
	return ( p.Page - 1 ) * p.Limit()
}