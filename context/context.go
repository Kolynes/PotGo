package context

import (
	"github.com/Kolynes/PotGo/controller"
	"github.com/Kolynes/PotGo/middleware"
)

type Context struct {
	routers    []*controller.Router
	Middleware []middleware.IMiddleware
}

func NewContext(routers []*controller.Router, middlewareInterfaces []middleware.IMiddleware) *Context {
	context := Context{
		routers:    routers,
		Middleware: middlewareInterfaces,
	}

	for index := 0; index < len(context.Middleware)-1; index++ {
		context.Middleware[index].SetNext(context.Middleware[index+1])
	}

	lastIndex := len(context.Middleware) - 1
	context.Middleware[lastIndex].SetNext(middleware.NewCommonMiddleware())
	return &context
}
