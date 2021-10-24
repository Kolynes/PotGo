package context

import (
	"github.com/Kolynes/PotGo/environment"
	"github.com/Kolynes/PotGo/types"
)

type Context struct {
	controllerContext *ControllerContext
	middlewareContext *MiddlewareContext
	Env               *environment.Environment
}

func (context *Context) GetControllerContext() types.IControllerContext {
	return context.controllerContext
}

func (context *Context) GetMiddlewareContext() types.IMiddlewareContext {
	return context.middlewareContext
}

func NewContext(env *environment.Environment) *Context {
	context := new(Context)
	context.controllerContext = NewControllerContext(context)
	context.middlewareContext = NewMiddlewareContext(context.controllerContext, context)
	context.Env = env
	return context
}
