package validations

import (
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
)

func GetDomainsValidator(req http.RequestPageContext) *helpers.HTTPError {
	if req.Authorizer == "" || req.Source == "" {
		return helpers.BadRequest("req authorizer or source should not be empty")
	}
	return nil
}

func GetUsersValidator(req http.RequestPageContext) *helpers.HTTPError {
	if req.Authorizer == "" || req.Source == "" {
		return helpers.BadRequest("req authorizer or source should not be empty")
	}
	return nil
}

func GenerateEmailValidator(req http.GenerateEmailRequest) *helpers.HTTPError {
	if req.RequestMeta.Authorizer == "" || req.RequestMeta.Source == "" {
		return helpers.BadRequest("req authorizer or source should not be empty")
	}
	if helpers.IsEmpty(req.Domain.Name) {
		return helpers.BadRequest("domain cannot be empty in generate email req")
	}
	if helpers.IsEmpty(req.Name.FirstName) || helpers.IsEmpty(req.Name.LastName) {
		return helpers.BadRequest("first/last name cannot be empty in generate email req")
	}

	// Later = more validations can be added based on email domain & first part length = 64 - overall length = 255
	return nil
}
