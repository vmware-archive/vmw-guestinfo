package rpcout

import (
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/sigma/vmw-guestinfo/message"
)

const RPCI_PROTOCOL_NUM int = 0x49435052

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

func (out *RpcOut) Send(request string) (string, error) {
	err := message.Send(out.channel, request)
	if err != nil {
		return "", err
	}

	reply, err := message.Receive(out.channel)
	if err != nil {
		return "", err
	}

	valid := true
	resp := ""
	if len(reply) < 2 {
		valid = false
	} else {
		resp = reply
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
