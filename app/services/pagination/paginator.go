package pagination

import (
	"github.com/jinzhu/gorm"
)

type Query struct {
	DB          *gorm.DB
	CurrentPage int
	PerPage     int
	OrderBy     []string
}

type Paginator struct {
	TotalRecords int         `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	Records      interface{} `json:"records"`
	Offset       int         `json:"offset"`
	PerPage      int         `json:"per_page"`
	CurrentPage  int         `json:"cur_page"`
	PrevPage     int         `json:"prev_page"`
	NextPage     int         `json:"next_page"`
}

/**Paginate given query and destination of model array returns paginator*/
func Paginate(p *Query, dest interface{}, scan bool) *Paginator {
	return nil
}
