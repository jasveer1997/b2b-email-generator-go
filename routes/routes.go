package handler

import (
	"context"
	"encoding/json"
	"fmt"
	http2 "github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/usecase"
	"github.com/jasveer1997/b2b-email-generator-go/utils"
	"net/http"
)

func GetDomainsHandler(ctx context.Context, usecaseImpl *usecase.UsecaseImpl) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqContext := utils.ReqContextQueryParser(r.URL.Query(), r.Header)
		res, errInRes := usecaseImpl.GetDomains(ctx, reqContext)
		if errInRes != nil {
			fmt.Println("Error Response from server:", errInRes.Error())
			http.Error(w, errInRes.Error(), (*errInRes).Status)
			return
		} else {
			fmt.Println("Response from server:", res)
			marshalledRes, errM := json.Marshal(res)
			if errM != nil {
				fmt.Println("Error marshalling Response from server:", errM.Error())
				http.Error(w, "Error marshalling Response from server", http.StatusInternalServerError)
				return
			} else {
				w.Write(marshalledRes)
			}
		}
	}
}

func GetUsersHandler(ctx context.Context, usecaseImpl *usecase.UsecaseImpl) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqContext := utils.ReqContextQueryParser(r.URL.Query(), r.Header)

		var requestData http2.GetUsersRequest
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}
		if requestData.Domains == nil {
			requestData.Domains = []string{}
		}

		res, errInRes := usecaseImpl.GetUsers(ctx, requestData, reqContext)
		if errInRes != nil {
			fmt.Println("Error Response from server:", errInRes.Error())
			http.Error(w, errInRes.Error(), (*errInRes).Status)
			return
		} else {
			fmt.Println("Response from server:", res)
			marshalledRes, errM := json.Marshal(res)
			if errM != nil {
				fmt.Println("Error marshalling Response from server:", errM.Error())
				http.Error(w, "Error marshalling Response from server", http.StatusInternalServerError)
				return
			} else {
				w.Write(marshalledRes)
			}
		}
	}
}

func GetGenerateEmailHandler(ctx context.Context, usecaseImpl *usecase.UsecaseImpl) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestData http2.GenerateEmailRequest
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		errInRes := usecaseImpl.GenerateEmail(ctx, requestData)
		if errInRes != nil {
			fmt.Println("Error Response from server:", errInRes.Error())
			http.Error(w, errInRes.Error(), (*errInRes).Status)
			return
		}
	}
}
