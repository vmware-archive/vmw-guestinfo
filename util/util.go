package util

import (
	"reflect"
	"runtime"
	"testing"
)

// Test utilities.

func AssertEqual(t *testing.T, a interface{}, b interface{}) bool {
	if !reflect.DeepEqual(a, b) {
		Fail(t)
		return false
	}

	return true
}

func AssertNoError(t *testing.T, err error) bool {
	if err != nil {
		t.Logf("error :%s", err.Error())
		Fail(t)
		return false
	}

	return true
}

func AssertNotNil(t *testing.T, a interface{}) bool {
	val := reflect.ValueOf(a)
	if val.IsNil() {
		Fail(t)
		return false
	}

	return true
}

func Fail(t *testing.T) {
	_, file, line, _ := runtime.Caller(2)
	t.Logf("FAIL on %s:%d", file, line)
	t.Fail()
}
