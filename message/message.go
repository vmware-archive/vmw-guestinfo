package message

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unsafe"
)

const (
	MESSAGE_TYPE_OPEN = iota
	MESSAGE_TYPE_SENDSIZE
	MESSAGE_TYPE_SENDPAYLOAD
	MESSAGE_TYPE_RECVSIZE
	MESSAGE_TYPE_RECVPAYLOAD
	MESSAGE_TYPE_RECVSTATUS
	MESSAGE_TYPE_CLOSE

	MESSAGE_STATUS_SUCCESS = uint16(0x0001)
	MESSAGE_STATUS_DORECV  = uint16(0x0002)
	MESSAGE_STATUS_CPT     = uint16(0x0010)
	MESSAGE_STATUS_HB      = uint16(0x0080)
)

var (
	// ErrChannelOpen represents a failure to open a channel
	ErrChannelOpen = errors.New("could not open channel")
	// ErrChannelClose represents a failure to close a channel
	ErrChannelClose = errors.New("could not close channel")
	// ErrRpciSend represents a failure to send a message
	ErrRpciSend = errors.New("unable to send RPCI command")
	// ErrRpciReceive represents a failure to receive a message
	ErrRpciReceive = errors.New("unable to receive RPCI command result")
)

type Channel struct {
	id uint16

	forcelowbandwidth bool
	buf               []byte

	cookie UInt64
}

// NewChannel opens a new Channel
func NewChannel(proto uint32) (*Channel, error) {
	flags := GUESTMSG_FLAG_COOKIE

retry:
	bp := &BackdoorProto{}

	bp.BX.Low.SetWord(proto | flags)
	bp.CX.Low.High = MESSAGE_TYPE_OPEN
	bp.CX.Low.Low = VMWARE_BDOOR_CMD_MESSAGE

	out := bp.InOut()
	if (out.CX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {
		if flags != 0 {
			flags = 0
			goto retry
		}

		Errorf("Message: Unable to open communication channel")
		return nil, ErrChannelOpen
	}

	ch := &Channel{}
	ch.id = out.DX.Low.High
	ch.cookie.High.SetWord(out.SI.Low.Word())
	ch.cookie.Low.SetWord(out.DI.Low.Word())

	Debugf("Opened channel %d", ch.id)
	return ch, nil
}

func (c *Channel) Close() error {
	bp := &BackdoorProto{}

	bp.CX.Low.High = MESSAGE_TYPE_CLOSE
	bp.CX.Low.Low = VMWARE_BDOOR_CMD_MESSAGE

	bp.DX.Low.High = c.id
	bp.SI.Low.SetWord(c.cookie.High.Word())
	bp.DI.Low.SetWord(c.cookie.Low.Word())

	out := bp.InOut()
	if (out.CX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {

		Errorf("Message: Unable to close communication channel %d", c.id)
		return ErrChannelClose
	}

	Debugf("Closed channel %d", c.id)
	return nil
}

func (c *Channel) Send(buf []byte) error {
retry:
	bp := &BackdoorProto{}
	bp.CX.Low.High = MESSAGE_TYPE_SENDSIZE
	bp.CX.Low.Low = VMWARE_BDOOR_CMD_MESSAGE

	bp.DX.Low.High = c.id
	bp.SI.Low.SetWord(c.cookie.High.Word())
	bp.DI.Low.SetWord(c.cookie.Low.Word())

	bp.BX.Low.SetWord(uint32(len(buf)))

	// send the size
	out := bp.InOut()
	if (out.CX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {
		Errorf("Message: Unable to send a message over the communication channel %d", c.id)
		return ErrRpciSend
	}

	if !c.forcelowbandwidth && (out.CX.Low.High&MESSAGE_STATUS_HB) == MESSAGE_STATUS_HB {
		hbbp := &BackdoorProto{}

		hbbp.BX.Low.Low = VMWARE_BDOORHB_CMD_MESSAGE
		hbbp.BX.Low.High = MESSAGE_STATUS_SUCCESS
		hbbp.DX.Low.High = c.id
		hbbp.BP.Low.SetWord(c.cookie.High.Word())
		hbbp.DI.Low.SetWord(c.cookie.Low.Word())
		hbbp.CX.Low.SetWord(uint32(len(buf)))
		hbbp.SI.SetQuad(uint64(uintptr(unsafe.Pointer(&buf[0]))))

		out := hbbp.HighBandwidthOut()
		if (out.BX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {

			if (out.BX.Low.High & MESSAGE_STATUS_CPT) != 0 {
				Debugf("A checkpoint occurred. Retrying the operation")
				goto retry
			}

			Errorf("Message: Unable to send a message over the communication channel %d", c.id)
			return ErrRpciSend
		}

	} else {

		bp.CX.Low.High = MESSAGE_TYPE_SENDPAYLOAD

		bbuf := bytes.NewBuffer(buf)
		for {

			// read 4 bytes at a time
			words := bbuf.Next(4)
			if len(words) == 0 {
				break
			}

			Debugf("sending %q over %d", string(words), c.id)
			switch len(words) {
			case 0:
				break

			case 3:
				bp.BX.Low.SetWord(binary.LittleEndian.Uint32([]byte{0x0, words[2], words[1], words[0]}))
			case 2:
				bp.BX.Low.SetWord(uint32(binary.LittleEndian.Uint16(words)))
			case 1:
				bp.BX.Low.SetWord(uint32(words[0]))

			default:
				bp.BX.Low.SetWord(binary.LittleEndian.Uint32(words))
			}

			out = bp.InOut()
			if (out.CX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {
				Errorf("Message: Unable to send a message over the communication channel %d", c.id)
				return ErrRpciSend
			}
		}
	}

	return nil
}

func (c *Channel) Receive() ([]byte, error) {
retry:
	var err error
	bp := &BackdoorProto{}
	bp.CX.Low.High = MESSAGE_TYPE_RECVSIZE
	bp.CX.Low.Low = VMWARE_BDOOR_CMD_MESSAGE

	bp.DX.Low.High = c.id
	bp.SI.Low.SetWord(c.cookie.High.Word())
	bp.DI.Low.SetWord(c.cookie.Low.Word())

	out := bp.InOut()
	if (out.CX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {
		Errorf("Message: Unable to poll for messages over the communication channel %d", c.id)
		return nil, ErrRpciReceive
	}

	if (out.CX.Low.High & MESSAGE_STATUS_DORECV) == 0 {
		Debugf("No message to retrieve")
		return nil, nil
	}

	// Receive the size.
	if out.DX.Low.High != MESSAGE_TYPE_SENDSIZE {
		Errorf("Message: Protocol error. Expected a MESSAGE_TYPE_SENDSIZE request from vmware")
		return nil, ErrRpciReceive
	}

	size := out.BX.Quad()
	var buf []byte

	if !c.forcelowbandwidth && (out.CX.Low.High&MESSAGE_STATUS_HB) == MESSAGE_STATUS_HB {
		buf = make([]byte, size)

		hbbp := &BackdoorProto{}

		hbbp.BX.Low.Low = VMWARE_BDOORHB_CMD_MESSAGE
		hbbp.BX.Low.High = MESSAGE_STATUS_SUCCESS
		hbbp.DX.Low.High = c.id
		hbbp.SI.Low.SetWord(c.cookie.High.Word())
		hbbp.BP.Low.SetWord(c.cookie.Low.Word())
		hbbp.CX.Low.SetWord(uint32(len(buf)))
		hbbp.DI.SetQuad(uint64(uintptr(unsafe.Pointer(&buf[0]))))

		out := hbbp.HighBandwidthIn()
		if (out.BX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {
			Errorf("Message: Unable to send a message over the communication channel %d", c.id)
			return nil, ErrRpciReceive
		}

	} else {

		b := bytes.NewBuffer(make([]byte, 0, size))

		for {

			if size == 0 {
				break
			}

			bp.CX.Low.High = MESSAGE_TYPE_RECVPAYLOAD
			bp.BX.Low.Low = MESSAGE_STATUS_SUCCESS

			out = bp.InOut()
			if (out.CX.Low.High & MESSAGE_STATUS_SUCCESS) == 0 {
				if (out.CX.Low.High & MESSAGE_STATUS_CPT) != 0 {
					Debugf("A checkpoint occurred. Retrying the operation")
					goto retry
				}

				Errorf("Message: Unable to receive a message over the communication channel %d", c.id)
				return nil, ErrRpciReceive
			}

			if out.DX.Low.High != MESSAGE_TYPE_SENDPAYLOAD {
				Errorf("Message: Protocol error. Expected a MESSAGE_TYPE_SENDPAYLOAD from vmware")
				return nil, ErrRpciReceive
			}

			Debugf("Received %#v", out.BX.Low.Word())

			switch size {
			case 1:
				err = binary.Write(b, binary.LittleEndian, uint8(out.BX.Low.Low))
				size = size - 1

			case 2:
				err = binary.Write(b, binary.LittleEndian, uint16(out.BX.Low.Low))
				size = size - 2

			case 3:
				err = binary.Write(b, binary.LittleEndian, uint16(out.BX.Low.Low))
				if err != nil {
					return nil, err
				}
				err = binary.Write(b, binary.LittleEndian, uint8(out.BX.Low.High))
				size = size - 3

			default:
				err = binary.Write(b, binary.LittleEndian, out.BX.Low.Word())
				size = size - 4
			}

			if err != nil {
				Errorf(err.Error())
				return nil, ErrRpciReceive
			}
		}

		buf = b.Bytes()
	}

	return buf, nil
}
