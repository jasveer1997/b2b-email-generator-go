package storage

import (
	"context"
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	"strings"
	"sync"
)

// Storage ideally holds connection to DB layers which can be utilised by implemented methods. For now, we are directly keeping a JSON of data here
type Storage struct {
	Domains    []StorageDomain
	domainLock *sync.RWMutex
	Users      []StorageUser
	userLock   *sync.RWMutex // This locks are kept to avoid parallel requests inconsistency (Write ahead)
}

func (storage *Storage) StoreUser(ctx context.Context, req StorageUser) *helpers.HTTPError {
	storage.userLock.Lock()
	currentUsers := storage.Users
	currentUsers = append(currentUsers, req)
	storage.Users = currentUsers
	storage.userLock.Unlock()
	return nil
}

func (storage *Storage) StoreDomain(ctx context.Context, req StorageDomain) *helpers.HTTPError {
	storage.domainLock.Lock()
	currentDomains := storage.Domains
	currentDomains = append(currentDomains, req)
	storage.Domains = currentDomains
	storage.domainLock.Unlock()
	return nil
}

func (storage *Storage) GetDomain(ctx context.Context, domain string) (StorageDomain, error) {
	storage.domainLock.RLock()
	currentDomains := storage.Domains
	// below loop is in memory - else it gets data from DB
	for _, currDomain := range currentDomains {
		if currDomain.DomainName == domain {
			return currDomain, nil
		}
	}
	storage.domainLock.RUnlock()
	//return nil, helpers.BadRequest("domain not found in DB, register it first?")
	return StorageDomain{
		DomainName: domain,
		EmailPref:  UNKNOWN,
	}, nil
}

func (storage *Storage) GetUser(ctx context.Context, req GetStorageUserRequest) StorageUser {
	storage.userLock.RLock()
	currentUsers := storage.Users
	// below loop is in memory - else it gets data from DB
	for _, currUser := range currentUsers {
		if currUser.FirstName == req.FirstName && currUser.LastName == req.LastName && currUser.Domain == req.Domain {
			return currUser
		}
	}
	storage.userLock.RUnlock()
	return StorageUser{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Domain:    req.Domain,
	}
}

func (storage *Storage) GetRegisteredUsers(ctx context.Context, req http.GetUsersRequest, reqMeta http.RequestPageContext) (*GetAllMatchingUsersResponse, *helpers.HTTPError) {
	storage.userLock.RLock()
	currentUsers := storage.Users
	matchingUsers := make([]StorageUser, 0)
	for _, currUser := range currentUsers {
		if len(req.Domains) != 0 {
			if helpers.ContainsMatchingDomain(req.Domains, currUser.Domain) {
				matchingUsers = append(matchingUsers, currUser)
			}
		} else {
			matchingUsers = append(matchingUsers, currUser)
		}
	}
	storage.userLock.RUnlock()
	if reqMeta.From > int32(len(matchingUsers)) {
		return nil, nil
	}
	to := reqMeta.From + reqMeta.Size
	if to > int32(len(matchingUsers)) {
		to = int32(len(matchingUsers))
	}
	return &GetAllMatchingUsersResponse{
		Users:      matchingUsers,
		Pagination: StoragePagination{Total: int32(len(matchingUsers))},
	}, nil
}

func (storage *Storage) GetAllMatchingDomains(ctx context.Context, reqMeta http.RequestPageContext) (*GetAllMatchingDomainsResponse, *helpers.HTTPError) {
	storage.domainLock.RLock()
	currentDomains := storage.Domains
	matchingDomains := make([]string, 0)
	for _, currDomain := range currentDomains {
		if reqMeta.Search != "" {
			if strings.Contains(currDomain.DomainName, reqMeta.Search) {
				matchingDomains = append(matchingDomains, currDomain.DomainName)
			}
		} else {
			matchingDomains = append(matchingDomains, currDomain.DomainName)
		}
	}
	storage.domainLock.RUnlock()
	if reqMeta.From > int32(len(matchingDomains)) {
		return nil, nil
	}
	to := reqMeta.From + reqMeta.Size
	if to > int32(len(matchingDomains)) {
		to = int32(len(matchingDomains))
	}
	return &GetAllMatchingDomainsResponse{
		Domains:    matchingDomains[reqMeta.From:to],
		Pagination: StoragePagination{Total: int32(len(matchingDomains))},
	}, nil
}

func GetStorage(ctx context.Context) (IStorage, *helpers.HTTPError) {
	// -- read static one time when server comes up and GetStorage() is called --
	initializedUserDataSetEqToDB := userData
	initializedDomainDataSetEqToDB := domainData

	return &Storage{
		Domains:    initializedDomainDataSetEqToDB,
		domainLock: &sync.RWMutex{},
		Users:      initializedUserDataSetEqToDB,
		userLock:   &sync.RWMutex{},
	}, nil
}
