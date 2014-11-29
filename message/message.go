package message

/*
#include <stdlib.h>
#include "message.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Channel *C.struct_Message_Channel

func Open(proto int) (Channel, error) {
	channel := C.Message_Open(C.uint32(proto))
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

func Send(c Channel, request []byte) error {
	buffer := (*C.uchar)(unsafe.Pointer(&request[0]))
	status := C.Message_Send(c, buffer, (C.size_t)(C.int(len(request)+1)))
	if status == 0 {
		return errors.New("Unable to send the RPCI command")
	}
	return nil
}

func Receive(c Channel) ([]byte, error) {
	var reply *C.uchar
	var reply_len C.size_t
	defer C.free(unsafe.Pointer(reply))

	status := C.Message_Receive(c, &reply, &reply_len)
	if status == 0 {
		return make([]byte, 0), errors.New("Unable to receive the result of the RPCI command")
	}

	return C.GoBytes(unsafe.Pointer(reply), (C.int)(reply_len)), nil
}
