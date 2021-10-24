package server

import (
	"testing"

	"github.com/Kolynes/PotGo/context"
	"github.com/Kolynes/PotGo/environment"
)

func TestServer(t *testing.T) {
	context := context.NewContext(&environment.Environment{})

	s := NewServer(
		context,
	)
	s.Initiate()
	s.baseInstance.Close()
}
