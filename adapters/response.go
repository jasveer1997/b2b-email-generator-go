package adapters

import (
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/external/storage"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
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

func AdaptGetUsersResponse(reqContext http.RequestPageContext, users *storage.GetAllMatchingUsersResponse) http.GetUsersResponse {
	res := make([]http.User, 0)
	for _, user := range users.Users {
		res = append(res, http.User{
			Name: http.FullName{
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			Domain: http.Domain{
				Name: user.Domain,
			},
			Email: user.Email,
		})
	}
	return http.GetUsersResponse{
		Users: res,
		Pagination: http.Pagination{
			From:  reqContext.From,
			Size:  reqContext.Size,
			Total: users.Pagination.Total,
		},
	}
}

func AdaptGetStorageUserRequest(req http.GenerateEmailRequest) storage.GetStorageUserRequest {
	firstName := req.Name.FirstName
	if !helpers.IsEmpty(req.Name.MiddleName) {
		firstName += " " + req.Name.MiddleName
	}
	return storage.GetStorageUserRequest{
		FirstName: req.Name.FirstName,
		LastName:  req.Name.LastName,
		Domain:    req.Domain.Name,
	}
}
