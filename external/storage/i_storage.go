package storage

import (
	"context"
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
)

type StorageUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Domain    string `json:"domain"`
	Email     string `json:"email"`
}

type StorageDomain struct {
	EmailPref  EmailPref `json:"email_pref"`
	DomainName string    `json:"domain"`
}

type StoragePagination struct {
	Total int32
}

type GetStorageUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Domain    string `json:"domain"`
}

type GetAllMatchingDomainsResponse struct {
	Domains    []string
	Pagination StoragePagination
}

type GetAllMatchingUsersResponse struct {
	Users      []StorageUser
	Pagination StoragePagination
}

type EmailPref string

const (
	UNKNOWN                      EmailPref = "UNKNOWN"
	FIRST_NAME_INITIAL_LAST_NAME EmailPref = "FIRST_NAME_INITIAL_LAST_NAME"
	FIRST_NAME_LAST_NAME         EmailPref = "FIRST_NAME_LAST_NAME"
)

type IStorage interface {
	GetAllMatchingDomains(ctx context.Context, reqMeta http.RequestPageContext) (*GetAllMatchingDomainsResponse, *helpers.HTTPError)
	GetRegisteredUsers(ctx context.Context, req http.GetUsersRequest, reqMeta http.RequestPageContext) (*GetAllMatchingUsersResponse, *helpers.HTTPError)
	GetUser(ctx context.Context, req GetStorageUserRequest) StorageUser
	GetDomain(ctx context.Context, domain string) (StorageDomain, error)
	StoreUser(ctx context.Context, req StorageUser) *helpers.HTTPError
	StoreDomain(ctx context.Context, req StorageDomain) *helpers.HTTPError
}
