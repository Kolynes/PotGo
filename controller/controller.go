package controller

import (
	"net/http"
	"reflect"

	"github.com/Kolynes/PotGo/response"
	"github.com/Kolynes/PotGo/types"
)

type Controller struct {
	route   string
	Context types.IContext
}

func (controller *Controller) SetContext(context types.IContext) {
	controller.Context = context
	controller.Context.GetControllerContext().RegisterController(controller)
}

func (controller *Controller) Route() string {
	return controller.route
}

func (c *Controller) CallRoute(controllerReflectionType reflect.Type, controllerReflectionValue reflect.Value, fragments []string, request *http.Request) (*response.Response, error) {
	if len(fragments) == 2 && fragments[1] != "CallRoute" {
		for i := 0; i < controllerReflectionType.NumMethod(); i++ {
			name := controllerReflectionType.Method(i).Name
			if name == fragments[1] {
				result := controllerReflectionValue.Method(i).Call([]reflect.Value{reflect.ValueOf(request)})
				err := result[1].Interface()
				if err == nil {
					return result[0].Interface().(*response.Response), nil
				} else {
					return result[0].Interface().(*response.Response), err.(error)
				}
			}
		}
	}
	notFoundResponse := new(response.Response)
	notFoundResponse.Status = 404
	notFoundResponse.Body = []byte("This resource could not be located")
	return notFoundResponse, nil
}
