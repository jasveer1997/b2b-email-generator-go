package utils

import (
	http2 "github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	"net/http"
	"net/url"
)

func ReqContextQueryParser(query url.Values, headers http.Header) http2.RequestPageContext {
	from := int32(0)
	if !helpers.IsEmpty(query.Get("from")) {
		from = helpers.ParseStrToInt32(query.Get("from"))
	}
	size := int32(10)
	if !helpers.IsEmpty(query.Get("size")) {
		size = helpers.ParseStrToInt32(query.Get("size"))
	}
	search := ""
	if !helpers.IsEmpty(query.Get("search")) {
		search = query.Get("search")
	}
	authorizer := ""
	if !helpers.IsEmpty(query.Get("authorizer")) {
		authorizer = query.Get("authorizer")
	}
	source := ""
	if !helpers.IsEmpty(query.Get("source")) {
		source = query.Get("source")
	}
	return http2.RequestPageContext{
		From:       from,
		Size:       size,
		Search:     search,
		Authorizer: authorizer,
		Source:     source,
	}
}
