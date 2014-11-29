package message

/*
#cgo LDFLAGS: /mnt/host/gocode/src/github.com/sigma/vmw-guestinfo/libs/libvmtools.a /usr/lib/x86_64-linux-gnu/libglib-2.0.a
#include <stdlib.h>
typedef struct Message_Channel Message_Channel;
extern Message_Channel* Message_Open(int proto);
extern int Message_Close(Message_Channel *chan);
extern int Message_Send(Message_Channel *chan, const char *buf, int bugSize);
extern int Message_Receive(Message_Channel *chan, char **buf, int *bugSize);
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Channel *C.struct_Message_Channel

func Open(proto int) (Channel, error) {
	channel := C.Message_Open(C.int(proto))
	if channel == nil {
		return nil, errors.New("Could not open channel")
	}
	return channel, nil
}

func Close(c Channel) error {
	status := C.Message_Close(c)
	if status == 0 {
		return errors.New("Could not close channel")
	}
	return nil
}

func Send(c Channel, request string) error {
	status := C.Message_Send(c, C.CString(request),
		C.int(len(request)+1))
	if status == 0 {
		return errors.New("Unable to send the RPCI command")
	}
	return nil
}

func Receive(c Channel) (string, error) {
	var reply *C.char
	var reply_len C.int
	defer C.free(unsafe.Pointer(reply))

	status := C.Message_Receive(c, &reply, &reply_len)
	if status == 0 {
		return "", errors.New("Unable to receive the result of the RPCI command")
	}

	return C.GoString(reply), nil
}
