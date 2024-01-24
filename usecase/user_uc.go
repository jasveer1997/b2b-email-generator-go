package usecase

import (
	"context"
	"github.com/jasveer1997/b2b-email-generator-go/adapters"
	"github.com/jasveer1997/b2b-email-generator-go/domain"
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/external/storage"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	"github.com/jasveer1997/b2b-email-generator-go/validations"
	"log"
)

type UsecaseImpl struct {
	Storage storage.IStorage
}

func (ucImpl *UsecaseImpl) GetDomains(ctx context.Context, reqContext http.RequestPageContext) (http.GetDomainsResponse, *helpers.HTTPError) {
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

	// 3 Return adapting the response
	return adapters.AdaptGetDomainsResponse(reqContext, domains), nil
}

func (ucImpl *UsecaseImpl) GetUsers(ctx context.Context, req http.GetUsersRequest, reqContext http.RequestPageContext) (http.GetUsersResponse, *helpers.HTTPError) {
	// 1 Validation
	err := validations.GetUsersValidator(reqContext)
	if err != nil {
		return http.GetUsersResponse{}, err
	}

	// 2 Read from storage directly, no need to introduce domain layer in read api
	users, err := ucImpl.Storage.GetRegisteredUsers(ctx, req, reqContext)
	if err != nil {
		return http.GetUsersResponse{}, err
	}

	// 3 Return adapting th response
	return adapters.AdaptGetUsersResponse(reqContext, users), nil
}

// GenerateEmail more preferably there should be a separate API to register domain preferences, rather than deriving it from data-set
func (ucImpl *UsecaseImpl) GenerateEmail(ctx context.Context, req http.GenerateEmailRequest) *helpers.HTTPError {
	// 1 basic Validation on req data
	err := validations.GenerateEmailValidator(req)
	if err != nil {
		return err
	}

	// 2 Read user from storage and convert to domain user
	storageUser := ucImpl.Storage.GetUser(ctx, adapters.AdaptGetStorageUserRequest(req))
	storageDomain, _ := ucImpl.Storage.GetDomain(ctx, req.Domain.Name) // IMP: ignoring error of domain not found - instead for now, let's register domain email pref based on first name - last name...
	domainUser := domain.GetDomainUserFromStorageUser(storageUser, storageDomain)

	if domainUser.UserExists {
		// if proper user already exists - at least in this API, do nothing (No updates!)
		log.Println("User already exists, no need to generate email")
		return nil
	}

	// start domain ops
	domainUser.GenerateEmail()

	// save to storage back!
	modifiedStorageUser := domain.GetStorageUserFromDomainUser(domainUser)
	errInStoringUser := ucImpl.Storage.StoreUser(ctx, modifiedStorageUser)
	if errInStoringUser != nil {
		log.Println("error occurred in domain user in storage")
		return errInStoringUser
	}

	if !domainUser.DomainExists {
		modifiedStorageDomain := domain.GetStorageDomainFromDomainUser(domainUser)
		errInStoringDomain := ucImpl.Storage.StoreDomain(ctx, modifiedStorageDomain)
		if errInStoringDomain != nil {
			log.Println("error occurred in domain in storage")
			return errInStoringDomain
		}
	}

	return nil
}

// GetNewUsecaseImpl initializes any downstream layer dependency for usage. Ex: Storage/Event layers
func GetNewUsecaseImpl(ctx context.Context) (*UsecaseImpl, *helpers.HTTPError) {
	storage, err := storage.GetStorage(ctx)
	if err != nil {
		return nil, err
	}
	return &UsecaseImpl{Storage: storage}, nil
}
