package utils

import (
	"testing"
)

func TestGetFuncName(t *testing.T) {
	fnName := "TestGetFuncName"
	name := GetFuncName(TestGetFuncName)
	if name != fnName {
		t.Errorf("expect %s but got %s", fnName, name)
	}
}
