package intacct

type ReadByQuery struct {
	Object   string `xml:"object"`
	Fields   string `xml:"fields"`
	Query    string `xml:"query"`
	PageSize int    `xml:"pagesize"`
}
