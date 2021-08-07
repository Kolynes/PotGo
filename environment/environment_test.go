package environment

import "testing"

func TestEnvironment(t *testing.T) {
	t.Log("testing environment")
	env := GetEnvironment()
	for key, value := range env.Variables {
		t.Log(key, value)
	}
}
