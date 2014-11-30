package message

import (
	"errors"

	"github.com/sigma/vmw-guestinfo/bridge"
)

type Channel bridge.MessageChannel

func Open(proto uint32) (Channel, error) {
	channel := bridge.MessageOpen(proto)
	if channel == nil {
		return nil, errors.New("Could not open channel")
	}
	return Channel(channel), nil
}

func Close(c Channel) error {
	status := bridge.MessageClose(bridge.MessageChannel(c))
	if !status {
		return errors.New("Could not close channel")
	}
	return nil
}

func Send(c Channel, request []byte) error {
	status := bridge.MessageSend(bridge.MessageChannel(c), request)
	if !status {
		return errors.New("Unable to send the RPCI command")
	}
	return nil
}

func Receive(c Channel) ([]byte, error) {
	res, status := bridge.MessageReceive(bridge.MessageChannel(c))
	if !status {
		return make([]byte, 0), errors.New("Unable to receive the result of the RPCI command")
	}
	return res, nil
}
