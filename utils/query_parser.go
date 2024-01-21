package utils

import (
	http2 "github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	"log"
	"net/http"
	"net/url"
)

func ReqContextQueryParser(query url.Values, headers http.Header) http2.RequestPageContext {
	from := int32(0)
	log.Println("from: ", query.Get("from"))
	log.Println("size: ", query.Get("size"))
	log.Println("search: ", query.Get("search"))
	log.Println("queries: ", query)
	if !helpers.IsEmpty(query.Get("from")) {
		from = helpers.ParseStrToInt32(query.Get("from"))
	}
	size := int32(0)
	if !helpers.IsEmpty(query.Get("size")) {
		size = helpers.ParseStrToInt32(query.Get("size"))
	}
	search := ""
	if !helpers.IsEmpty(query.Get("search")) {
		search = query.Get("search")
	}
	authorizer := ""
	if !helpers.IsEmpty(headers.Get("authorizer")) {
		authorizer = headers.Get("authorizer")
	}
	source := ""
	if !helpers.IsEmpty(headers.Get("source")) {
		source = headers.Get("source")
	}
	return http2.RequestPageContext{
		From:       from,
		Size:       size,
		Search:     search,
		Authorizer: authorizer,
		Source:     source,
	}
}
