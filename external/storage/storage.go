package storage

import (
	"encoding/json"
	"github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	gee "github.com/tbxark/g4vercel"
	"log"
	"os"
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
	var initializedUserDataSetEqToDB []StorageUser
	var initializedDomainDataSetEqToDB []StorageDomain

	// -- read json one time when server comes up and GetStorage() is called --
	// user
	jsonData, err := os.ReadFile("external/storage/static_initial_user_data.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
		return nil, helpers.InternalServerError("User JSON reading failed. See logs for more detail")
	}
	err = json.Unmarshal(jsonData, &initializedUserDataSetEqToDB)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
		return nil, helpers.InternalServerError("User JSON unmarshalling failed. See logs for more detail")
	}
	// domain
	jsonData, err = os.ReadFile("external/storage/static_initial_domain_data.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
		return nil, helpers.InternalServerError("Domain JSON reading failed. See logs for more detail")
	}
	err = json.Unmarshal(jsonData, &initializedDomainDataSetEqToDB)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
		return nil, helpers.InternalServerError("Domain JSON unmarshalling failed. See logs for more detail")
	}

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
