package http

import (
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	gee "github.com/tbxark/g4vercel"
)

type IUsecases interface {
	GetDomains(ctx *gee.Context, reqContext RequestPageContext) (GetDomainsResponse, *helpers.HTTPError)
	GetUsers(ctx *gee.Context, reqContext RequestPageContext) (GetUsersResponse, *helpers.HTTPError)
	GenerateEmail(ctx *gee.Context, req GenerateEmailRequest) (User, *helpers.HTTPError)
}

type Pagination struct {
	From  int32 `json:"from"`
	Size  int32 `json:"size"`
	Total int32 `json:"total"`
}

type RequestPageContext struct {
	From       int32  `json:"from"`
	Size       int32  `json:"size"`
	Search     string `json:"search"`
	Authorizer string `json:"authorizer"`
	Source     string `json:"source"`
}

type GetDomainsResponse struct {
	Domains    []Domain   `json:"domains"`
	Pagination Pagination `json:"pagination"`
}

type GetUsersResponse struct {
	Users      []User     `json:"users"`
	Pagination Pagination `json:"pagination"`
}

// RequestMeta is a basic meta details to track actor information
type RequestMeta struct {
	Authorizer string `json:"authorizer"`
	Source     string `json:"source"`
}

type Domain struct { // Can be separated based on further use-cases for request, responses
	Name string `json:"name"`
	// RegisteredEmployees int32  `json:"registered_employees"` // optional later implemetation
}

type FullName struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type User struct {
	Name   FullName `json:"name"`
	Domain Domain   `json:"domain"`
	Email  string   `json:"email"`
}

type GenerateEmailRequest struct {
	Name        FullName    `json:"full_name"`
	Domain      Domain      `json:"domain"`
	RequestMeta RequestMeta `json:"request_meta"`
}
