package rpcout

import (
	"errors"
	"fmt"

	"github.com/sigma/vmw-guestinfo/message"
)

var ErrRpciFormat = errors.New("invalid format for RPCI command result")

const RPCI_PROTOCOL_NUM uint32 = 0x49435052

func SendOne(format string, a ...interface{}) (reply []byte, ok bool, err error) {
	request := fmt.Sprintf(format, a...)
	return SendOneRaw([]byte(request))
}

func SendOneRaw(request []byte) (reply []byte, ok bool, err error) {
	out := &RpcOut{}
	if err = out.Start(); err != nil {
		return
	}
	if reply, ok, err = out.Send(request); err != nil {
		return
	}
	if err = out.Stop(); err != nil {
		return
	}
	return
}

type RpcOut struct {
	channel message.Channel
}

func (out *RpcOut) Start() error {
	channel, err := message.Open(RPCI_PROTOCOL_NUM)
	if err != nil {
		return err
	}
	out.channel = channel
	return nil
}

func (out *RpcOut) Stop() error {
	err := message.Close(out.channel)
	out.channel = nil
	return err
}

func (out *RpcOut) Send(request []byte) (reply []byte, ok bool, err error) {
	if err = message.Send(out.channel, request); err != nil {
		return
	}

	var resp []byte
	if resp, err = message.Receive(out.channel); err != nil {
		return
	}

	switch string(resp[:2]) {
	case "0 ":
		reply = resp[2:]
	case "1 ":
		reply = resp[2:]
		ok = true
	default:
		err = ErrRpciFormat
	}
	return
}
