package server

import (
	"testing"

	"github.com/Kolynes/PotGo/context"
	"github.com/Kolynes/PotGo/controller"
	"github.com/Kolynes/PotGo/middleware"
)

func TestServer(t *testing.T) {
	s := Server{
		Host: "127.0.0.1",
		Port: 8000,
	}
	s.Initiate(
		context.NewContext([]*controller.Router{}, []middleware.IMiddleware{}),
	)
}
