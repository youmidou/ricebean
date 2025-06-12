package _ymd

import (
	"bytes"

	"ricebean/pkg/conn/packet"
)

// PomeloPacketDecoder reads and decodes network data slice following pomelo's protocol
type PomeloPacketDecoder struct{}

// NewPomeloPacketDecoder returns a new decoder that used for decode network bytes slice.
func NewPomeloPacketDecoder() *PomeloPacketDecoder {
	return &PomeloPacketDecoder{}
}

func (c *PomeloPacketDecoder) forward(buf *bytes.Buffer) (int, packet.Type, error) {
	header := buf.Next(HeadLength)
	return ParseHeader(header)
}

// Decode decode the network bytes slice to packet.Packet(s)
func (c *PomeloPacketDecoder) Decode(data []byte) ([]*packet.Packet, error) {
	buf := bytes.NewBuffer(nil)
	buf.Write(data)

	var (
		packets []*packet.Packet
		err     error
	)
	// check length
	if buf.Len() < HeadLength {
		return nil, nil
	}

	// first time
	size, typ, err := c.forward(buf)
	if err != nil {
		return nil, err
	}

	for size <= buf.Len() {
		p := &packet.Packet{Type: typ, Length: size, Data: buf.Next(size)}
		packets = append(packets, p)

		// if no more packets, break
		if buf.Len() < HeadLength {
			break
		}

		size, typ, err = c.forward(buf)
		if err != nil {
			return nil, err
		}
	}

	return packets, nil
}
