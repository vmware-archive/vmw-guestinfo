package message

import "github.com/vmware/vmw-guestinfo/bdoor"

func HypervisorPortCheck() bool {
	p := &bdoor.BackdoorProto{}

	p.CX.Low.SetWord(bdoor.CommandGetVersion)

	out := p.InOut()

	Infof("version %d", out.AX.Low.Word())
	return 0 != out.AX.Low.Word()
}
