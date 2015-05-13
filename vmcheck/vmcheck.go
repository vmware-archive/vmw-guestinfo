package vmcheck

import (
	"github.com/sigma/vmw-guestinfo/bridge"
)

func IsVirtualWorld() bool {
	return bridge.VmCheckIsVirtualWorld()
}

func GetVersion() (version uint32, typ uint32) {
	return bridge.VmCheckGetVersion()
}
