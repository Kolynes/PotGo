package middleware

import (
	"net/http"

	"github.com/Kolynes/PotGo/response"
)

type Middleware struct {
	next           IMiddleware
	HandleRequest  func(*http.Request)
	HandleResponse func(*response.Response)
}

type IMiddleware interface {
	Handle(*http.Request) (*response.Response, error)
	SetNext(IMiddleware)
	GetNext() IMiddleware
}

func (middleware *Middleware) SetNext(next IMiddleware) {
	middleware.next = next
}

func (middleware *Middleware) GetNext() IMiddleware {
	return middleware.next
}

func (middleware *Middleware) Handle(request *http.Request) (*response.Response, error) {
	middleware.HandleRequest(request)
	response, err := middleware.next.Handle(request)
	middleware.HandleResponse(response)
	return response, err
}
