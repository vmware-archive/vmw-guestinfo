package vmcheck

import (
	"github.com/sigma/vmw-guestinfo/bridge"
)

func IsVirtualWorld() bool {
	return bridge.VmCheckIsVirtualWorld()
}

func GetVersion() (uint32, uint32) {
	return bridge.VmCheckGetVersion()
}
