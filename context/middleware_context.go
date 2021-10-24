package context

import (
	"net/http"

	"github.com/Kolynes/PotGo/response"
	"github.com/Kolynes/PotGo/types"
)

type MiddlewareContext struct {
	middleware        []types.IMiddleware
	controllerContext types.IMiddleware
	parentContext     types.IContext
}

func (context *MiddlewareContext) Handle(request *http.Request) *response.Response {
	return context.middleware[0].Handle(request)
}

func (context *MiddlewareContext) RegisterMiddleware(middleware types.IMiddleware) {
	if len(context.middleware) > 0 {
		context.middleware[len(context.middleware)-1].SetNext(middleware)
	}
	middleware.SetNext(context.controllerContext)
	middleware.SetContext(context.parentContext)
	context.middleware = append(context.middleware, middleware)
}

func NewMiddlewareContext(controllerContext types.IMiddleware, parentContext types.IContext) *MiddlewareContext {
	context := new(MiddlewareContext)
	context.middleware = make([]types.IMiddleware, 0)
	context.controllerContext = controllerContext
	context.parentContext = parentContext
	return context
}
