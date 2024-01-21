package adapters

import (
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/external/storage"
)

func AdaptGetDomainsResponse(reqContext http.RequestPageContext, domains *storage.GetAllMatchingDomainsResponse) http.GetDomainsResponse {
	res := make([]http.Domain, 0)
	for _, domain := range domains.Domains {
		res = append(res, http.Domain{Name: domain})
	}
	return http.GetDomainsResponse{
		Domains: res,
		Pagination: http.Pagination{
			From:  reqContext.From,
			Size:  reqContext.Size,
			Total: domains.Pagination.Total,
		},
	}
}
