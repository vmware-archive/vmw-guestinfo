package message

import (
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
