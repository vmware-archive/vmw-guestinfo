package bdoor

type UInt32 struct {
	High uint16
	Low  uint16
}

func (u *UInt32) Word() uint32 {
	return uint32(u.High)<<16 + uint32(u.Low)
}

func (u *UInt32) SetWord(w uint32) {
	u.High = uint16(w >> 16)
	u.Low = uint16(w)
}

type UInt64 struct {
	High UInt32
	Low  UInt32
}

func (u *UInt64) Quad() uint64 {
	return uint64(u.High.Word())<<32 + uint64(u.Low.Word())
}

func (u *UInt64) SetQuad(w uint64) {
	u.High.SetWord(uint32(w >> 32))
	u.Low.SetWord(uint32(w))
}
