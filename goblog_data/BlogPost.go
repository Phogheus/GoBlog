package goblog_data

import "time"

type BlogPost struct {
	Id              int
	Author          string
	DatePosted      time.Time
	DateLastUpdated time.Time
	Title           string
	Body            string
}
