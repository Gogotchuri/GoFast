package pagination

import (
	"math"

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
	db := p.DB
	//Paginator to return
	paginator := &Paginator{}

	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	//Order database
	for _, o := range p.OrderBy {
		db = db.Order(o)
	}
	//Get total records count
	if scan {
		db.Count(&paginator.TotalRecords)
	} else {
		db.Model(dest).Count(&paginator.TotalRecords)
	}
	//Calculate offset
	paginator.Offset = (p.CurrentPage - 1) * p.PerPage
	//Write into destination and append array elements to paginator records
	if scan {
		db.Limit(p.PerPage).Offset(paginator.Offset).Scan(dest)
	} else {
		db.Limit(p.PerPage).Offset(paginator.Offset).Find(dest)
	}
	paginator.Records = dest
	paginator.CurrentPage = p.CurrentPage
	paginator.PerPage = p.PerPage
	paginator.TotalPages = int(math.Ceil(float64(paginator.TotalRecords) / float64(p.PerPage)))
	//Set previous page
	paginator.PrevPage = p.CurrentPage - 1
	if p.CurrentPage == 1 {
		paginator.PrevPage = p.CurrentPage
	}
	//Set next page
	paginator.NextPage = p.CurrentPage + 1
	if p.CurrentPage == paginator.TotalPages {
		paginator.NextPage = p.CurrentPage
	}
	return paginator
}
