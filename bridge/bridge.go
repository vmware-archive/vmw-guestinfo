package bridge

/*
#cgo CFLAGS: -I../include
#include <stdlib.h>
#include "message.h"
#include "vmcheck.h"
void Warning(const char *fmt, ...) {}
void Debug(const char *fmt, ...) {}
void Panic(const char *fmt, ...) {}
void Log(const char *fmt, ...) {}
*/
import "C"
import "unsafe"

type MessageChannel *C.struct_Message_Channel

func MessageOpen(proto uint32) MessageChannel {
	return C.Message_Open(C.uint32(proto))
}

func MessageClose(c MessageChannel) bool {
	status := C.Message_Close(c)
	return status != 0
}

func MessageSend(c MessageChannel, request []byte) bool {
	buffer := (*C.uchar)(unsafe.Pointer(&request[0]))
	status := C.Message_Send(c, buffer, (C.size_t)(C.int(len(request)+1)))
	return status != 0
}

func MessageReceive(c MessageChannel) ([]byte, bool) {
	var reply *C.uchar
	var reply_len C.size_t
	defer C.free(unsafe.Pointer(reply))

	status := C.Message_Receive(c, &reply, &reply_len)

	res := C.GoBytes(unsafe.Pointer(reply), (C.int)(reply_len))
	return res, status != 0
}

func VmCheckIsVirtualWorld() bool {
	return C.VmCheck_IsVirtualWorld() != 0
}

func VmCheckGetVersion() (uint32, uint32) {
	var version C.uint32
	var typ C.uint32
	C.VmCheck_GetVersion(&version, &typ)
	return uint32(version), uint32(typ)
}
