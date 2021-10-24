package context

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/Kolynes/PotGo/response"
	"github.com/Kolynes/PotGo/types"
)

type ControllerContext struct {
	parentContext types.IContext
	controllers   []types.IController
}

func (context *ControllerContext) Handle(request *http.Request) *response.Response {
	fragments := strings.Split(request.URL.Path, "/")
	if len(fragments) > 0 {
		for _, controller := range context.controllers {
			route := controller.Route()
			if route == fragments[0] {
				controllerType := reflect.TypeOf(&controller)
				controllerValue := reflect.ValueOf(&controller)
				newResponse, err := controller.CallRoute(controllerType, controllerValue, fragments, request)
				if err != nil {
					return &response.Response{
						Status: 500,
						Body:   []byte(err.Error()),
					}
				} else {
					return newResponse
				}
			}
		}
	}
	return &response.Response{
		Status: 404,
		Body:   []byte("Resource Not Found"),
	}
}

func (context *ControllerContext) SetNext(middleware types.IMiddleware) {}

func (context *ControllerContext) SetContext(parentContext types.IContext) {
	context.parentContext = parentContext
}

func (context *ControllerContext) RegisterController(c types.IController) {
	c.SetContext(context.parentContext)
	context.controllers = append(context.controllers, c)
}

func NewControllerContext(parentContext types.IContext) *ControllerContext {
	context := new(ControllerContext)
	context.controllers = []types.IController{}
	context.parentContext = parentContext
	return context
}
