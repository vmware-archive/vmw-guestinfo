package main

import (
	"fmt"
	"unsafe"
)

// #cgo LDFLAGS: libs/libvmtools.a /usr/lib/x86_64-linux-gnu/libglib-2.0.a
/*
extern char * RpcVMX_ConfigGetString(const char *defval, const char *var);
#include <stdlib.h>
*/
import "C"

func main() {
	cs := C.RpcVMX_ConfigGetString(C.CString("foo"), C.CString("foo"))
	defer C.free(unsafe.Pointer(cs))

	s := C.GoString(cs)
	fmt.Println(s)
}
