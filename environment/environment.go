package environment

import (
	"os"
	"strings"
)

type Environment struct {
	Variables map[string]string
}

type IEnvironment interface {
	Get(key string) interface{}
	Setup()
}

func (env *Environment) Get(key string) string {
	return env.Variables[key]
}

func GetEnvironment() *Environment {
	env := Environment{}
	env.Variables = make(map[string]string)
	for _, e := range os.Environ() {
		keyValue := strings.SplitN(e, "=", 2)
		env.Variables[keyValue[0]] = keyValue[1]
	}
	return &env
}
