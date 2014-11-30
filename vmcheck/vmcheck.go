package vmcheck

import (
	"github.com/sigma/vmw-guestinfo/bridge"
)

func IsVirtualWorld() bool {
	return bridge.VmCheckIsVirtualWorld()
}
