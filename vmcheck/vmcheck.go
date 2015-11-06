package vmcheck

import (
	"github.com/vmware/vmw-guestinfo/bridge"
)

// IsVirtualWorld returns whether the code is running in a VMware virtual machine or no
func IsVirtualWorld() bool {
	return bridge.VMCheckIsVirtualWorld()
}

// GetVersion returns the hypervisor version
func GetVersion() (version uint32, typ uint32) {
	return bridge.VMCheckGetVersion()
}
