package middleware

import (
	"net/http"

	"github.com/Kolynes/PotGo/response"
	"github.com/Kolynes/PotGo/types"
)

type Middleware struct {
	Next    types.IMiddleware
	Context types.IContext
}

func (middleware *Middleware) SetNext(next types.IMiddleware) {
	middleware.Next = next
}

func (middleware *Middleware) Handle(request *http.Request) *response.Response {
	return new(response.Response)
}

func (middleware *Middleware) SetContext(context types.IContext) {
	middleware.Context = context
	middleware.Context.GetMiddlewareContext().RegisterMiddleware(middleware)
}
