package storage

import (
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	gee "github.com/tbxark/g4vercel"
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

type GetAllMatchingDomainsResponse struct {
	Domains    []string
	Pagination StoragePagination
}

type EmailPref string

const (
	FIRST_NAME_INITIAL_LAST_NAME EmailPref = "FIRST_NAME_INITIAL_LAST_NAME"
	FIRST_NAME_LAST_NAME         EmailPref = "FIRST_NAME_LAST_NAME"
)

type IStorage interface {
	GetAllMatchingDomains(ctx *gee.Context, reqMeta http.RequestPageContext) (*GetAllMatchingDomainsResponse, *helpers.HTTPError)
}
