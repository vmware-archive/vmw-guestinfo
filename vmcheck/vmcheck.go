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

package vmcheck

import "github.com/vmware/vmw-guestinfo/bdoor"

// IsVirtualWorld returns true if running in a VM.
// NOTE:  This will PANIC if not run in a virtual world
func IsVirtualWorld() bool {
	p := &bdoor.BackdoorProto{}
	p.CX.Low.SetWord(bdoor.CommandGetVersion)

	// TODO(FA) get this inside a fork() call and collect the return code since
	// we can't mask the SIGSEGV signal.
	out := p.InOut()

	return 0 != out.AX.Low.Word()
}
