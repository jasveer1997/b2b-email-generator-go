package domain

import (
	"github.com/jasveer1997/b2b-email-generator-go/external/storage"
	"github.com/jasveer1997/b2b-email-generator-go/helpers"
	"strings"
)

type Employee struct {
	FirstName    string
	MiddleName   string
	LastName     string
	Domain       string
	Email        string
	DomainFormat storage.EmailPref
	UserExists   bool
	DomainExists bool
}

func GetStorageUserFromDomainUser(user Employee) storage.StorageUser {
	firstName := user.FirstName
	if !helpers.IsEmpty(user.MiddleName) {
		firstName += " " + user.MiddleName
	}
	return storage.StorageUser{
		FirstName: firstName,
		LastName:  user.LastName,
		Domain:    user.Domain,
		Email:     user.Email,
	}
}

// GetStorageDomainFromDomainUser is kept directly for now on Employee to reduce code - it should be handled separately though!
func GetStorageDomainFromDomainUser(user Employee) storage.StorageDomain {
	return storage.StorageDomain{
		EmailPref:  user.DomainFormat,
		DomainName: user.Domain,
	}
}

func GetDomainUserFromStorageUser(user storage.StorageUser, domainData storage.StorageDomain) Employee {
	return Employee{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Domain:       user.Domain,
		Email:        user.Email, // put whatever is there from storage (empty, if it does not exists - new user!)
		DomainFormat: domainData.EmailPref,
		UserExists:   !helpers.IsEmpty(user.Email),
		DomainExists: domainData.EmailPref != storage.UNKNOWN,
	}
}

func (dUser *Employee) GenerateEmail() {
	switch dUser.DomainFormat {
	case storage.FIRST_NAME_INITIAL_LAST_NAME:
		firstChar := strings.ToLower(dUser.FirstName[:1])
		dUser.Email = firstChar + strings.ToLower(dUser.LastName) + "@" + dUser.Domain
	case storage.FIRST_NAME_LAST_NAME:
		dUser.Email = strings.ToLower(dUser.FirstName) + strings.ToLower(dUser.LastName) + "@" + dUser.Domain // Imp: Ignoring Middle_name!
	case storage.UNKNOWN: // UNKNOWN will follow FIRST_NAME_LAST_NAME as default behavior
		dUser.Email = strings.ToLower(dUser.FirstName) + strings.ToLower(dUser.LastName) + "@" + dUser.Domain // Imp: Ignoring Middle_name!
		dUser.DomainFormat = storage.FIRST_NAME_LAST_NAME
		// Ideally we should attach an event here as well - like DOMAIN_ADDED (more ideally should go to diff DOMAIN server apis though)/DOMAIN_FORMAT_CREATED
	}
}
