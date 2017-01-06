package message

const (
	VMWARE_BDOOR_MAGIC  = uint64(0x564D5868)
	VMWARE_BDOOR_PORT   = uint16(0x5658)
	VMWARE_BDOORHB_PORT = uint16(0x5659)

	VMWARE_BDOOR_CMD_GETVERSION           = uint32(10)
	VMWARE_BDOOR_CMD_GETDEVICELISTELEMENT = uint32(11)
	VMWARE_BDOOR_CMD_TOGGLEDEVICE         = uint32(12)

	VMWARE_BDOOR_CMD_MESSAGE   = uint16(0x1e)
	VMWARE_BDOORHB_CMD_MESSAGE = uint16(0)
	GUESTMSG_FLAG_COOKIE       = uint32(0x80000000)
)

type BackdoorProto struct {
	// typedef union {
	//   struct {
	//      DECLARE_REG_NAMED_STRUCT(ax);
	//      size_t size; /* Register bx. */
	//      DECLARE_REG_NAMED_STRUCT(cx);
	//      DECLARE_REG_NAMED_STRUCT(dx);
	//      DECLARE_REG_NAMED_STRUCT(si);
	//      DECLARE_REG_NAMED_STRUCT(di);
	//   } in;
	//   struct {
	//      DECLARE_REG_NAMED_STRUCT(ax);
	//      DECLARE_REG_NAMED_STRUCT(bx);
	//      DECLARE_REG_NAMED_STRUCT(cx);
	//      DECLARE_REG_NAMED_STRUCT(dx);
	//      DECLARE_REG_NAMED_STRUCT(si);
	//      DECLARE_REG_NAMED_STRUCT(di);
	//   } out;
	// } proto;

	AX, BX, CX, DX, SI, DI, BP UInt64
	size                       uint32
}

func bdoor_inout(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64)
func bdoor_hbout(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64)
func bdoor_hbin(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64)
func bdoor_inout_test(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64)

func (p *BackdoorProto) InOut() *BackdoorProto {
	p.DX.Low.Low = VMWARE_BDOOR_PORT
	p.AX.SetQuad(VMWARE_BDOOR_MAGIC)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_inout(
		p.AX.Quad(),
		p.BX.Quad(),
		p.CX.Quad(),
		p.DX.Quad(),
		p.SI.Quad(),
		p.DI.Quad(),
		p.BP.Quad(),
	)

	ret := &BackdoorProto{}
	ret.AX.SetQuad(retax)
	ret.BX.SetQuad(retbx)
	ret.CX.SetQuad(retcx)
	ret.DX.SetQuad(retdx)
	ret.SI.SetQuad(retsi)
	ret.DI.SetQuad(retdi)
	ret.BP.SetQuad(retbp)

	return ret
}

func (p *BackdoorProto) HighBandwidthOut() *BackdoorProto {

	p.DX.Low.Low = VMWARE_BDOORHB_PORT
	p.AX.SetQuad(VMWARE_BDOOR_MAGIC)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_hbout(
		p.AX.Quad(),
		p.BX.Quad(),
		p.CX.Quad(),
		p.DX.Quad(),
		p.SI.Quad(),
		p.DI.Quad(),
		p.BP.Quad(),
	)

	ret := &BackdoorProto{}
	ret.AX.SetQuad(retax)
	ret.BX.SetQuad(retbx)
	ret.CX.SetQuad(retcx)
	ret.DX.SetQuad(retdx)
	ret.SI.SetQuad(retsi)
	ret.DI.SetQuad(retdi)
	ret.BP.SetQuad(retbp)

	return ret
}

func (p *BackdoorProto) HighBandwidthIn() *BackdoorProto {

	p.DX.Low.Low = VMWARE_BDOORHB_PORT
	p.AX.SetQuad(VMWARE_BDOOR_MAGIC)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_hbin(
		p.AX.Quad(),
		p.BX.Quad(),
		p.CX.Quad(),
		p.DX.Quad(),
		p.SI.Quad(),
		p.DI.Quad(),
		p.BP.Quad(),
	)

	ret := &BackdoorProto{}
	ret.AX.SetQuad(retax)
	ret.BX.SetQuad(retbx)
	ret.CX.SetQuad(retcx)
	ret.DX.SetQuad(retdx)
	ret.SI.SetQuad(retsi)
	ret.DI.SetQuad(retdi)
	ret.BP.SetQuad(retbp)

	return ret
}
