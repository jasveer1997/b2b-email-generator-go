package tests

import (
	"context"
	http2 "github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/external/storage"
	"github.com/jasveer1997/b2b-email-generator-go/usecase"
	"log"
	"testing"
)

func TestEmailGenerator(t *testing.T) {
	// Create an instance of the mock
	ctx := context.Background()
	usecaseImpl, _ := usecase.GetNewUsecaseImpl(ctx)

	err := usecaseImpl.GenerateEmail(ctx, http2.GenerateEmailRequest{
		Name: http2.FullName{
			FirstName: "test",
			LastName:  "user",
		},
		Domain: http2.Domain{
			Name: "babbel.com",
		},
		RequestMeta: http2.RequestMeta{
			Authorizer: "test",
			Source:     "test",
		},
	})

	if err != nil {
		t.Errorf("error executing case")
	}

	// defining expected result -
	expectedUser := storage.StorageUser{
		FirstName: "test",
		LastName:  "user",
		Domain:    "babbel.com",
		Email:     "tuser@babbel.com",
	}

	storageUser := usecaseImpl.Storage.GetUser(ctx, storage.GetStorageUserRequest{
		FirstName: "test",
		LastName:  "user",
		Domain:    "babbel.com",
	})

	log.Println(storageUser)

	// Assert user with email
	if storageUser != expectedUser {
		t.Errorf("Expected: %v result, got %v", expectedUser, storageUser)
	}
}

// No need to actual mock on any layer - as all layers are in memory only
