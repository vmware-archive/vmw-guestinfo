package message

import (
	"reflect"
	"runtime"
	"testing"

	sigma "github.com/sigma/bdoor"
)

func TestHypervisorPortCheck(t *testing.T) {
	if !sigma.HypervisorPortCheck() {
		t.Skip("Not in a virtual world")
		return
	}

	t.Log("Running in a VM: ", HypervisorPortCheck())
}

func TestBdoorArgAlignment(t *testing.T) {
	a := uint64(0xFFFFFFFF0000022)
	b := uint64(33)
	c := uint64(44)
	d := uint64(55)
	si := uint64(0xFFFFFFFF0000066)
	di := uint64(0xFFFAAFFF0000077)
	bp := uint64(0xFFFFFFFFAAAAAAA)

	oa, ob, oc, od, osi, odi, obp := bdoor_inout_test(a, b, c, d, si, di, bp)

	if !AssertEqual(t, a, oa) ||
		!AssertEqual(t, b, ob) ||
		!AssertEqual(t, c, oc) ||
		!AssertEqual(t, d, od) ||
		!AssertEqual(t, si, osi) ||
		!AssertEqual(t, di, odi) ||
		!AssertEqual(t, bp, obp) {
		return
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) bool {
	if !reflect.DeepEqual(a, b) {
		Fail(t)
		return false
	}

	return true
}

func AssertNoError(t *testing.T, err error) bool {
	if err != nil {
		t.Log("error :%s", err.Error())
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
