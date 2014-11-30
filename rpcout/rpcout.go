package rpcout

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/sigma/vmw-guestinfo/message"
)

const RPCI_PROTOCOL_NUM uint32 = 0x49435052

func SendOne(format string, a ...interface{}) ([]byte, error) {
	request := fmt.Sprintf(format, a...)
	return SendOneRaw([]byte(request))
}

func SendOneRaw(request []byte) ([]byte, error) {
	var reply []byte
	glog.Infof("Rpci: Sending request='%s'", request)

	status := false

	out := &RpcOut{}
	err := out.Start()
	if err != nil {
		reply = []byte("RpcOut: Unable to open the communication channel")
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
	channel message.Channel
}

func (out *RpcOut) Start() error {
	channel, err := message.Open(RPCI_PROTOCOL_NUM)
	if err != nil {
		return errors.New("Could not open channel with RPCI protocol")
	}
	out.channel = channel
	return nil
}

func (out *RpcOut) Stop() error {
	err := message.Close(out.channel)
	out.channel = nil
	return err
}

func (out *RpcOut) Send(request []byte) ([]byte, error) {
	resp := make([]byte, 0)
	err := message.Send(out.channel, request)
	if err != nil {
		return resp, err
	}

	reply, err := message.Receive(out.channel)
	if err != nil {
		return resp, err
	}

	valid := true
	if len(reply) < 2 {
		valid = false
	} else {
		resp = reply
		prefix := resp[:2]

		if bytes.Equal(prefix, []byte("1 ")) {
			resp = resp[2:]
		} else if bytes.Equal(prefix, []byte("0 ")) {
			resp = make([]byte, 0)
		} else {
			valid = false
		}
	}

	if valid {
		return resp, nil
	} else {
		return nil, errors.New("Invalid format for the result of the RPCI command")
	}
}
