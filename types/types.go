package types

import (
	"net/http"
	"reflect"

	"github.com/Kolynes/PotGo/response"
)

type IContext interface {
	GetControllerContext() IControllerContext
	GetMiddlewareContext() IMiddlewareContext
}

type IMiddleware interface {
	Handle(*http.Request) *response.Response
	SetNext(IMiddleware)
	SetContext(context IContext)
}

type IControllerContext interface {
	IMiddleware
	RegisterController(IController)
}

type IMiddlewareContext interface {
	Handle(*http.Request) *response.Response
	RegisterMiddleware(middleware IMiddleware)
}

type IController interface {
	CallRoute(reflect.Type, reflect.Value, []string, *http.Request) (*response.Response, error)
	SetContext(context IContext)
	Route() string
}

type IWebsocketController interface {
	Upgrade(writer http.ResponseWriter, request *http.Request) (bool, error)
	reader()
	onMessage()
	onClose()
}

type IServer interface {
	http.Handler
	Initiate()
}
