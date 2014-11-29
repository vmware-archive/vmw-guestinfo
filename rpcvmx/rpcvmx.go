package rpcvmx

import (
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: /mnt/host/gocode/src/github.com/sigma/vmw-guestinfo/libs/libvmtools.a /usr/lib/x86_64-linux-gnu/libglib-2.0.a
extern int RpcOut_sendOne(char **reply, int *repLen, char const *reqFmt);
#include <stdlib.h>
*/
import "C"

func ConfigGetString(key string, default_value string) string {
	cmd := fmt.Sprintf("info-get guestinfo.%s", key)

	var val *C.char
	defer C.free(unsafe.Pointer(val))

	if C.RpcOut_sendOne(&val, nil, C.CString(cmd)) == 0 {
		return default_value
	}
	return C.GoString(val)

}
