package paging

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	// DefaultLimit ...
	DefaultLimit = "10"
)

// Paging ...
type Paging struct {
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Count    int64       `json:"count"`
	Records  interface{} `json:"records"`
}

// NewPaging creates a paging instance
func NewPaging(records interface{}, offset int, limit int, count int64) *Paging {
	return &Paging{
		Offset:  offset,
		Limit:   limit,
		Count:   count,
		Records: records,
	}
}

// Init initialize Next and Previous fields
func (me *Paging) Init(request *http.Request) {
	var (
		limit, offset int
	)

	u := request.URL
	q := u.Query()

	if q.Get("limit") == "" {
		limit, _ = strconv.Atoi(DefaultLimit)
	} else {
		limit, _ = strconv.Atoi(q.Get("limit"))
	}

	if q.Get("offset") == "" {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(q.Get("offset"))
	}

	if me.Count > int64(me.Limit)+int64(me.Offset) {
		q.Set("offset", strconv.Itoa(offset+limit))
		u.RawQuery, _ = url.QueryUnescape(q.Encode())
		me.Next = u.String()
	}

	if me.Offset > 0 {
		q.Set("offset", strconv.Itoa(offset-limit))
		u.RawQuery, _ = url.QueryUnescape(q.Encode())
		me.Previous = u.String()
	}

	return
}
