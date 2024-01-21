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
