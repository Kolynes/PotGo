package context

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/Kolynes/PotGo/controller"
	"github.com/Kolynes/PotGo/environment"
	"github.com/Kolynes/PotGo/middleware"
	"github.com/Kolynes/PotGo/response"
)

type TestMiddleware struct {
	middleware.Middleware
}

func (testMiddleware *TestMiddleware) Handle(request *http.Request) *response.Response {
	return &response.Response{}
}

type TestController struct {
	controller.Controller
}

func (c *TestController) HelloWorld(request *http.Request) *response.Response {
	return &response.Response{
		Status: 200,
	}
}

func TestNewMiddlewareContext(t *testing.T) {
	env := environment.Environment{}
	parentContext := NewContext(&env)
	controllerContext := NewControllerContext(parentContext)
	context := NewMiddlewareContext(controllerContext, parentContext)
	t.Log(context)
}

func TestRegisterMiddleware(t *testing.T) {
	middleware := TestMiddleware{}
	env := environment.Environment{}
	parentContext := NewContext(&env)
	controllerContext := NewControllerContext(parentContext)
	context := NewMiddlewareContext(controllerContext, parentContext)
	context.RegisterMiddleware(&middleware)
}

func TestMiddlewareLinking(t *testing.T) {
	middleware := TestMiddleware{}
	env := environment.Environment{}
	parentContext := NewContext(&env)
	controllerContext := NewControllerContext(parentContext)
	context := NewMiddlewareContext(controllerContext, parentContext)
	context.RegisterMiddleware(&middleware)
	middleware1 := TestMiddleware{}
	context.RegisterMiddleware(&middleware1)
	if middleware.Next != &middleware1 {
		t.Error("Linking failed")
	}
}

func TestNewControllerContext(t *testing.T) {
	env := environment.Environment{}
	parentContext := NewContext(&env)
	controllerContext := NewControllerContext(parentContext)
	t.Log(controllerContext)
}

func TestRegisterController(t *testing.T) {
	controller := TestController{}
	env := environment.Environment{}
	parentContext := NewContext(&env)
	context := NewControllerContext(parentContext)
	context.RegisterController(&controller)
	if len(context.controllers) != 1 {
		t.Error("Failed to register controller")
	}
}

func TestControllerContextHandle404(t *testing.T) {
	controller := TestController{}
	env := environment.Environment{}
	parentContext := NewContext(&env)
	context := NewControllerContext(parentContext)
	context.RegisterController(&controller)
	if len(context.controllers) != 1 {
		t.Error("Failed to register controller")
	}
	request := http.Request{
		URL: &url.URL{
			Path: "TesController/HelloWorld",
		},
	}
	response := context.Handle(&request)
	t.Log(response)
}

func TestControllerContextHandle(t *testing.T) {
	controller := TestController{}
	env := environment.Environment{}
	parentContext := NewContext(&env)
	context := NewControllerContext(parentContext)
	context.RegisterController(&controller)
	if len(context.controllers) != 1 {
		t.Error("Failed to register controller")
	}
	request := http.Request{
		URL: &url.URL{
			Path: "TestController/HelloWorld",
		},
	}
	response := context.Handle(&request)
	t.Log(response)
}
