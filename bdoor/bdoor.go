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
	BackdoorPort       = uint16(0x5658)
	BackdoorHighBWPort = uint16(0x5659)

	CommandGetVersion = uint32(10)

	CommandMessage       = uint16(0x1e)
	CommandHighBWMessage = uint16(0)
	CommandFlagCookie    = uint32(0x80000000)
)

func (p *BackdoorProto) InOut() *BackdoorProto {
	p.DX.SetShort(BackdoorPort)
	setReg(p.AX, BackdoorMagic)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_inout(
		getReg(p.AX),
		getReg(p.BX),
		getReg(p.CX),
		getReg(p.DX),
		getReg(p.SI),
		getReg(p.DI),
		getReg(p.BP),
	)

	ret := &BackdoorProto{}
	setReg(ret.AX, retax)
	setReg(ret.BX, retbx)
	setReg(ret.CX, retcx)
	setReg(ret.DX, retdx)
	setReg(ret.SI, retsi)
	setReg(ret.DI, retdi)
	setReg(ret.BP, retbp)

	return ret
}

func (p *BackdoorProto) HighBandwidthOut() *BackdoorProto {
	p.DX.SetShort(BackdoorHighBWPort)
	setReg(p.AX, BackdoorMagic)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_hbout(
		getReg(p.AX),
		getReg(p.BX),
		getReg(p.CX),
		getReg(p.DX),
		getReg(p.SI),
		getReg(p.DI),
		getReg(p.BP),
	)

	ret := &BackdoorProto{}
	setReg(ret.AX, retax)
	setReg(ret.BX, retbx)
	setReg(ret.CX, retcx)
	setReg(ret.DX, retdx)
	setReg(ret.SI, retsi)
	setReg(ret.DI, retdi)
	setReg(ret.BP, retbp)

	return ret
}

func (p *BackdoorProto) HighBandwidthIn() *BackdoorProto {
	p.DX.SetShort(BackdoorHighBWPort)
	setReg(p.AX, BackdoorMagic)

	retax, retbx, retcx, retdx, retsi, retdi, retbp := bdoor_hbin(
		getReg(p.AX),
		getReg(p.BX),
		getReg(p.CX),
		getReg(p.DX),
		getReg(p.SI),
		getReg(p.DI),
		getReg(p.BP),
	)

	ret := &BackdoorProto{}
	setReg(ret.AX, retax)
	setReg(ret.BX, retbx)
	setReg(ret.CX, retcx)
	setReg(ret.DX, retdx)
	setReg(ret.SI, retsi)
	setReg(ret.DI, retdi)
	setReg(ret.BP, retbp)

	return ret
}
