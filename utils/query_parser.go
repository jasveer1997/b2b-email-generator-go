package utils

import (
	http2 "github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	"net/http"
	"net/url"
)

func ReqContextQueryParser(query url.Values, headers http.Header) http2.RequestPageContext {
	from := int32(0)
	if helpers.IsEmpty(query["from"]) {
		from = helpers.ParseStrToInt32(query["from"][0])
	}
	size := int32(0)
	if helpers.IsEmpty(query["size"]) {
		size = helpers.ParseStrToInt32(query["size"][0])
	}
	search := ""
	if helpers.IsEmpty(query["size"]) {
		search = query["search"][0]
	}
	authorizer := ""
	if helpers.IsEmpty(headers["authorizer"]) {
		authorizer = headers["authorizer"][0]
	}
	source := ""
	if helpers.IsEmpty(headers["source"]) {
		source = headers["source"][0]
	}
	return http2.RequestPageContext{
		From:       from,
		Size:       size,
		Search:     search,
		Authorizer: authorizer,
		Source:     source,
	}
}
