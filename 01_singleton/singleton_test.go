package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 == ins2 {
		t.Log("instance is equal")
	}
}
