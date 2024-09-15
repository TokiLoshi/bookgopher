package gobookapi

import (
	"time"
)

type ArticlesToRead struct {
	Title string 
	Author string
	Genre string
	URL string
	DateAdded time.Time
}

type ReadArticles struct {
	Title string
	Author string
	Genre string
	URL string 
	DateAdded time.Time 
	Notes string
}