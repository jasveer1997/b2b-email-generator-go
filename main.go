package main

import (
	"encoding/json"
	"fmt"
	http2 "github.com/jasveer1997/b2b-email-generator-go/external/http"
	"github.com/jasveer1997/b2b-email-generator-go/usecase"
	"net/http"
)

func main() {
	usecaseImpl, err := usecase.GetNewUsecaseImpl()
	if err != nil {
		panic(err.Error())
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		//query := r.URL.Query()
		//headers := r.Header
		// reqContext := utils.ReqContextQueryParser(query, headers)

		res, err := usecaseImpl.GetDomains(nil, http2.RequestPageContext{
			From:       0,
			Size:       10,
			Search:     "",
			Authorizer: "someone",
			Source:     "b2b-ui-app",
		})
		if err != nil {
			fmt.Println("Error Response from server:", err.Error())
		} else {
			fmt.Println("Response from server:", res)
			marshalledRes, errM := json.Marshal(res)
			if errM != nil {
				fmt.Println("Error marshalling Response from server:", errM.Error())
			} else {
				w.Write(marshalledRes)
			}
		}
	}

	http.HandleFunc("/domains", handler)

	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err2 := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err2 != nil {
		fmt.Println("Error starting server:", err2)
	}
}
