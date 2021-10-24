package controller

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/Kolynes/PotGo/response"
)

type TestController struct {
	Controller
}

func (c *TestController) HelloWorldError(request *http.Request) (*response.Response, error) {
	return &response.Response{
		Status: 500,
	}, errors.New("sorry")
}

func (c *TestController) HelloWorld(request *http.Request) (*response.Response, error) {
	return &response.Response{
		Status: 200,
	}, nil
}

func TestCallRoute(t *testing.T) {
	controller := TestController{}
	request := &http.Request{
		URL: &url.URL{
			Path: "TestController/HelloWorld",
		},
	}
	controller.CallRoute(reflect.TypeOf(&controller), reflect.ValueOf(&controller), []string{"TestController", "HelloWorld"}, request)
}

func TestCallRouteError(t *testing.T) {
	controller := TestController{}
	request := &http.Request{
		URL: &url.URL{
			Path: "TestController/HelloWorldError",
		},
	}
	controller.CallRoute(reflect.TypeOf(&controller), reflect.ValueOf(&controller), []string{"TestController", "HelloWorldError"}, request)
}

func TestCallRoute404(t *testing.T) {
	controller := TestController{}
	request := &http.Request{
		URL: &url.URL{
			Path: "TestController/Helloorld",
		},
	}
	controller.CallRoute(reflect.TypeOf(&controller), reflect.ValueOf(&controller), []string{"TestController", "Helloorld"}, request)
}
