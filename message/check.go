package message

func HypervisorPortCheck() bool {
	p := &BackdoorProto{}

	p.CX.Low.SetWord(VMWARE_BDOOR_CMD_GETVERSION)

	out := p.InOut()

	Infof("version %d", out.AX.Low.Word())
	return 0 != out.AX.Low.Word()
}
