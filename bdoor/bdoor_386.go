// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bdoor

const (
	BackdoorMagic      = uint32(0x564D5868)
	BackdoorPort       = uint16(0x5658)
	BackdoorHighBWPort = uint16(0x5659)

	CommandGetVersion = uint32(10)

	CommandMessage       = uint16(0x1e)
	CommandHighBWMessage = uint16(0)
	CommandFlagCookie    = uint32(0x80000000)
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

	AX, BX, CX, DX, SI, DI, BP UInt32
	size                       uint32
}

func bdoor_inout(ax, bx, cx, dx, si, di, bp uint32) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint32)
func bdoor_hbout(ax, bx, cx, dx, si, di, bp uint32) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint32)
func bdoor_hbin(ax, bx, cx, dx, si, di, bp uint32) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint32)
func bdoor_inout_test(ax, bx, cx, dx, si, di, bp uint32) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint32)

func (p *BackdoorProto) InOut() *BackdoorProto {
	p.DX.Low = BackdoorPort
	p.AX.SetWord(BackdoorMagic)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_inout(
		p.AX.Word(),
		p.BX.Word(),
		p.CX.Word(),
		p.DX.Word(),
		p.SI.Word(),
		p.DI.Word(),
		p.BP.Word(),
	)

	ret := &BackdoorProto{}
	ret.AX.SetWord(retax)
	ret.BX.SetWord(retbx)
	ret.CX.SetWord(retcx)
	ret.DX.SetWord(retdx)
	ret.SI.SetWord(retsi)
	ret.DI.SetWord(retdi)
	ret.BP.SetWord(retbp)

	return ret
}

func (p *BackdoorProto) HighBandwidthOut() *BackdoorProto {
	p.DX.Low = BackdoorHighBWPort
	p.AX.SetWord(BackdoorMagic)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_hbout(
		p.AX.Word(),
		p.BX.Word(),
		p.CX.Word(),
		p.DX.Word(),
		p.SI.Word(),
		p.DI.Word(),
		p.BP.Word(),
	)

	ret := &BackdoorProto{}
	ret.AX.SetWord(retax)
	ret.BX.SetWord(retbx)
	ret.CX.SetWord(retcx)
	ret.DX.SetWord(retdx)
	ret.SI.SetWord(retsi)
	ret.DI.SetWord(retdi)
	ret.BP.SetWord(retbp)

	return ret
}

func (p *BackdoorProto) HighBandwidthIn() *BackdoorProto {
	p.DX.Low = BackdoorHighBWPort
	p.AX.SetWord(BackdoorMagic)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_hbin(
		p.AX.Word(),
		p.BX.Word(),
		p.CX.Word(),
		p.DX.Word(),
		p.SI.Word(),
		p.DI.Word(),
		p.BP.Word(),
	)

	ret := &BackdoorProto{}
	ret.AX.SetWord(retax)
	ret.BX.SetWord(retbx)
	ret.CX.SetWord(retcx)
	ret.DX.SetWord(retdx)
	ret.SI.SetWord(retsi)
	ret.DI.SetWord(retdi)
	ret.BP.SetWord(retbp)

	return ret
}
