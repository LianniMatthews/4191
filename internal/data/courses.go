package data

// School represents one row of data in database
type Course struct {
	Code   string `json:"course"`
	Title  string `json:"title"`
	Credit int64  `json:"credit"`
}
