package duckduckgo

import (
	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/yields/ant"
)

var page struct {
	Title string `css:"title"`
}

func SearchThis(query string) *paginator.Paginator {
	// what's ctx supposed to look like? apparently something from context package https://pkg.go.dev/context#pkg-overview
	page, _ := ant.Fetch(ctx, "duckduckgo.com/?q="+query)
}
