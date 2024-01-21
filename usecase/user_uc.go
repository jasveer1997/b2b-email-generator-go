package usecase

import (
	"github.com/jasveer1997/b2b-email-generator-go/adapters"
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/external/storage"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	"github.com/jasveer1997/b2b-email-generator-go/validations"
	gee "github.com/tbxark/g4vercel"
)

type UsecaseImpl struct {
	Storage storage.IStorage
}

func (ucImpl *UsecaseImpl) GetDomains(ctx *gee.Context, reqContext http.RequestPageContext) (http.GetDomainsResponse, *helpers.HTTPError) {
	// 1 Validation
	err := validations.GetDomainsValidator(reqContext)
	if err != nil {
		return http.GetDomainsResponse{}, err
	}

	// 2 Read from storage directly, no need to introduce domain layer in read api
	domains, err := ucImpl.Storage.GetAllMatchingDomains(ctx, reqContext)
	if err != nil {
		return http.GetDomainsResponse{}, nil
	}

	// 3 Return adapting th response
	return adapters.AdaptGetDomainsResponse(reqContext, domains), nil
}

func (ucImpl *UsecaseImpl) GetUsers(ctx *gee.Context, reqContext http.RequestPageContext) (http.GetUsersResponse, *helpers.HTTPError) {
	return http.GetUsersResponse{}, nil
}

// more preferrably there should be a separate API to register domain preferences, rather than deriving it from data-set
func (ucImpl *UsecaseImpl) GenerateEmail(ctx *gee.Context, req http.GenerateEmailRequest) (http.User, *helpers.HTTPError) {
	return http.User{}, nil
}

// GetNewUsecaseImpl initializes any downstream layer dependency for usage. Ex: Storage/Event layers
func GetNewUsecaseImpl() (*UsecaseImpl, *helpers.HTTPError) {
	storage, err := storage.GetStorage()
	if err != nil {
		return nil, err
	}
	return &UsecaseImpl{Storage: storage}, nil
}
