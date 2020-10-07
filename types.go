package intacct

type ReadByQuery struct {
	Object   string `xml:"object"`
	Fields   string `xml:"fields"`
	Query    Query  `xml:"query"`
	PageSize int    `xml:"pagesize"`
}

type Query interface {
	Query() string
}

// func (q *Query) AddFilter(filter Filter) {
// }

// func (q *Query) AddCondition(condition Condition) {
// }

// type Filter struct {
// }

// type Condition struct {
// }

type NoQuery struct{}

func (r NoQuery) Query() string {
	return ""
}

type RawQuery string

func (r RawQuery) Query() string {
	return string(r)
}
