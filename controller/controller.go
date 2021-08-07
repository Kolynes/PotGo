package controller

import (
	"errors"
	"net/http"
	"reflect"
	"runtime"

	"github.com/Kolynes/PotGo/response"
)

type Controller struct {
	Route string
}

type IController interface {
	Handle(routeFragment string, request *http.Request) (response.Response, error)
}

func (this *Controller) Handle(routeFragment string, request *http.Request) (response.Response, error) {
	var null response.Response
	if routeFragment == "Handle" {
		return null, errors.New("invalid route fragment 'handle'")
	}
	controllerReflection := reflect.ValueOf(this)
	for i := 0; i < controllerReflection.NumMethod(); i++ {
		if runtime.FuncForPC(controllerReflection.Method(i).Pointer()).Name() == routeFragment {
			result := controllerReflection.Method(i).Call([]reflect.Value{reflect.ValueOf(request)})
			return result[0].Interface().(response.Response), result[1].Interface().(error)
		}
	}
	return null, errors.New("no matching route found")
}
