package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jasveer1997/b2b-email-generator-go/routes"
	"github.com/jasveer1997/b2b-email-generator-go/usecase"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// initialize mux and ctx, cors (cors needed for separate FE, BE domains)
	r := mux.NewRouter()
	ctx := context.Background()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})

	// initialize usecase layer
	usecaseImpl, err := usecase.GetNewUsecaseImpl(ctx)
	if err != nil {
		panic(err.Error())
	}

	// initialize route handlers
	domainHandler := handler.GetDomainsHandler(ctx, usecaseImpl)
	userHandler := handler.GetUsersHandler(ctx, usecaseImpl)
	generateEmailHandler := handler.GetGenerateEmailHandler(ctx, usecaseImpl)

	// routes definition
	r.HandleFunc("/domains", domainHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", userHandler).Methods(http.MethodPost)
	r.HandleFunc("/generate_email", generateEmailHandler).Methods(http.MethodPut)

	// up the server
	errInServer := http.ListenAndServe(":8081", corsHandler.Handler(r))
	if errInServer != nil {
		log.Fatalf("Failed to start server. Error: ", errInServer)
	}
}
