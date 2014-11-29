package rpcout

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/golang/glog"
)

/*
#cgo LDFLAGS: /mnt/host/gocode/src/github.com/sigma/vmw-guestinfo/libs/libvmtools.a /usr/lib/x86_64-linux-gnu/libglib-2.0.a
#include <stdlib.h>
typedef struct Message_Channel Message_Channel;
extern Message_Channel* Message_Open(int proto);
extern int Message_Close(Message_Channel *chan);
extern int Message_Send(Message_Channel *chan, const char *buf, int bugSize);
extern int Message_Receive(Message_Channel *chan, char **buf, int *bugSize);
int RPCI_PROTOCOL_NUM = 0x49435052;
*/
import "C"

func SendOne(format string, a ...interface{}) (string, error) {
	request := fmt.Sprintf(format, a...)
	return SendOneRaw(request)
}

func SendOneRaw(request string) (string, error) {
	var reply string
	glog.Infof("Rpci: Sending request='%s'", request)

	status := false

	out := &RpcOut{}
	err := out.Start()
	if err != nil {
		reply = "RpcOut: Unable to open the communication channel"
	} else {
		reply, err = out.Send(request)
		if err != nil {
			status = true
		}
	}

	glog.Infof("Rpci: Sent request='%s', reply='%s', status=%t",
		request, reply, status)

	err = out.Stop()
	if err != nil {
		glog.Infof("Rpci: unable to close the communication channel")
	}

	return reply, err
}

type RpcOut struct {
	channel *C.struct_Message_Channel
}

func (out *RpcOut) Start() error {
	channel := C.Message_Open(C.RPCI_PROTOCOL_NUM)
	if channel == nil {
		return errors.New("could not open channel with RPCI protocol")
	}
	out.channel = channel
	return nil
}

func (out *RpcOut) Stop() error {
	status := C.Message_Close(out.channel)
	if status == 0 {
		return errors.New("could not close channel")
	}
	out.channel = nil
	return nil
}

func (out *RpcOut) Send(request string) (string, error) {
	status := C.Message_Send(out.channel, C.CString(request),
		C.int(len(request)+1))
	if status == 0 {
		return "", errors.New("Unable to send the RPCI command")
	}

	var reply *C.char
	var reply_len C.int
	defer C.free(unsafe.Pointer(reply))
	status = C.Message_Receive(out.channel, &reply, &reply_len)
	if status == 0 {
		return "", errors.New("Unable to receive the result of the RPCI command")
	}

	valid := true
	resp := ""
	if reply_len < 2 {
		valid = false
	} else {
		resp = C.GoString(reply)
		prefix := resp[:2]

		if prefix == "1 " {
			resp = resp[2:]
		} else if prefix == "0 " {
			resp = ""
		} else {
			valid = false
		}
	}

	if valid {
		return resp, nil
	} else {
		return "", errors.New("Invalid format for the result of the RPCI command")
	}
}
