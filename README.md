# paging
--
    import "github.com/frozentech/paging"


## Usage

```go
const (
	// DefaultLimit ...
	DefaultLimit = "10"
)
```

#### type Paging

```go
type Paging struct {
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Count    int64       `json:"count"`
	Records  interface{} `json:"records"`
}
```

Paging ...

#### func  NewPaging

```go
func NewPaging(records interface{}, offset int, limit int, count int64) *Paging
```
NewPaging creates a paging instance

#### func (*Paging) Init

```go
func (me *Paging) Init(request *http.Request)
```
Init initialize Next and Previous fields
