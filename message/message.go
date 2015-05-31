package message

import (
	"errors"

	"github.com/sigma/vmw-guestinfo/bridge"
)

var (
	ErrChannelOpen  = errors.New("could not open channel")
	ErrChannelClose = errors.New("could not close channel")
	ErrRpciSend     = errors.New("unable to send RPCI command")
	ErrRpciReceive  = errors.New("unable to receive RPCI command result")
)

type Channel bridge.MessageChannel

func Open(proto uint32) (Channel, error) {
	if channel := bridge.MessageOpen(proto); channel != nil {
		return Channel(channel), nil
	}
	return nil, ErrChannelOpen
}

func Close(c Channel) error {
	if status := bridge.MessageClose(bridge.MessageChannel(c)); status {
		return nil
	}
	return ErrChannelClose
}

func Send(c Channel, request []byte) error {
	if status := bridge.MessageSend(bridge.MessageChannel(c), request); status {
		return nil
	}
	return ErrRpciSend
}

func Receive(c Channel) ([]byte, error) {
	if res, status := bridge.MessageReceive(bridge.MessageChannel(c)); status {
		return res, nil
	}
	return nil, ErrRpciReceive
}
