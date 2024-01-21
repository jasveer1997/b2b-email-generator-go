package storage

import (
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	gee "github.com/tbxark/g4vercel"
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

func GetStorage() (IStorage, *helpers.HTTPError) {
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

func (storage *Storage) GetAllMatchingDomains(ctx *gee.Context, reqMeta http.RequestPageContext) (*GetAllMatchingDomainsResponse, *helpers.HTTPError) {
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
