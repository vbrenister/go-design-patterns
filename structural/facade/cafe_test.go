package facade

import "testing"


func TestMakeAmericano(t *testing.T) {
	makeAmericano(100)
}

func TestMakeLatte(t *testing.T) {
	makeLatte(100, true)
}