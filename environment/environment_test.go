package environment

import (
	"testing"
)

func TestEnvironment(t *testing.T) {
	t.Log("testing environment")
	env := GetEnvironment(map[string]interface{}{})
	for key, value := range *env {
		t.Log(key, value)
	}
}
