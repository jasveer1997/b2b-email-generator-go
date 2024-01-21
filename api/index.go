package handler

import (
	"fmt"
	"github.com/jasveer1997/b2b-email-generator-go/usecase"
	"github.com/jasveer1997/b2b-email-generator-go/utils"
	. "github.com/tbxark/g4vercel"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	usecaseImpl, err := usecase.GetNewUsecaseImpl()
	if err != nil {
		panic(err.Error())
	}
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))
	server.GET("/domains", func(context *Context) {
		query := context.Req.URL.Query()
		headers := context.Req.Header
		reqContext := utils.ReqContextQueryParser(query, headers)
		res, err := usecaseImpl.GetDomains(context, reqContext)
		if err != nil {
			context.JSON(err.Status, H{
				"message": err.Message,
			})
		} else {
			context.JSON(200, H{
				"response": res,
			})
		}
	})
	server.Handle(w, r)
}
