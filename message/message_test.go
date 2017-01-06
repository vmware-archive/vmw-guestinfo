package message

import (
	"testing"

	sigma "github.com/sigma/bdoor"
	"github.com/vmware/vmw-guestinfo/util"
)

const rpciProtocolNum uint32 = 0x49435052

func TestOpenClose(t *testing.T) {
	l := DefaultLogger.(*logger)
	l.DebugLevel = true

	if !sigma.HypervisorPortCheck() {
		t.Skip("Not in a virtual world")
		return
	}

	ch, err := NewChannel(rpciProtocolNum)
	if !util.AssertNotNil(t, ch) || !util.AssertNoError(t, err) {
		return
	}

	// check low bandwidth
	ch.forceLowBW = true
	err = ch.Send([]byte("info-get guestinfo.doesnotexistdoesnotexit"))
	if !util.AssertNoError(t, err) {
		return
	}

	b, err := ch.Receive()
	if !util.AssertNoError(t, err) || !util.AssertNotNil(t, b) {
		return
	}

	if !util.AssertEqual(t, "0 No value found", string(b)) {
		return
	}

	if !util.AssertNoError(t, ch.Close()) {
		return
	}

	// check high bandwidth
	ch, err = NewChannel(rpciProtocolNum)
	if !util.AssertNotNil(t, ch) || !util.AssertNoError(t, err) {
		return
	}

	err = ch.Send([]byte("info-get guestinfo.doesnotexistdoesnotexit"))
	if !util.AssertNoError(t, err) {
		return
	}

	b, err = ch.Receive()
	if !util.AssertNoError(t, err) || !util.AssertNotNil(t, b) {
		return
	}

	if !util.AssertEqual(t, "0 No value found", string(b)) {
		return
	}

	if !util.AssertNoError(t, ch.Close()) {
		return
	}
}
