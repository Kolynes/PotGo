package environment

import (
	"os"
	"strings"
)

type Environment map[string]interface{}

func GetEnvironment(userSettings map[string]interface{}) *Environment {
	env := Environment{}
	for _, e := range os.Environ() {
		keyValue := strings.SplitN(e, "=", 2)
		env[keyValue[0]] = keyValue[1]
	}
	for key, value := range userSettings {
		env[key] = value
	}
	return &env
}
