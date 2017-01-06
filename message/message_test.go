package message

import (
	"testing"

	sigma "github.com/sigma/bdoor"
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
	if !AssertNotNil(t, ch) || !AssertNoError(t, err) {
		return
	}

	// check low bandwidth
	ch.forcelowbandwidth = true
	err = ch.Send([]byte("info-get guestinfo.doesnotexistdoesnotexit"))
	if !AssertNoError(t, err) {
		return
	}

	b, err := ch.Receive()
	if !AssertNoError(t, err) || !AssertNotNil(t, b) {
		return
	}

	if !AssertEqual(t, "0 No value found", string(b)) {
		return
	}

	if !AssertNoError(t, ch.Close()) {
		return
	}

	// check high bandwidth
	ch, err = NewChannel(rpciProtocolNum)
	if !AssertNotNil(t, ch) || !AssertNoError(t, err) {
		return
	}

	err = ch.Send([]byte("info-get guestinfo.doesnotexistdoesnotexit"))
	if !AssertNoError(t, err) {
		return
	}

	b, err = ch.Receive()
	if !AssertNoError(t, err) || !AssertNotNil(t, b) {
		return
	}

	if !AssertEqual(t, "0 No value found", string(b)) {
		return
	}

	if !AssertNoError(t, ch.Close()) {
		return
	}
}
