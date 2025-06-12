// Copyright (c) nano Author and TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package codec

import (
	"bytes"

	"ricebean/pkg/conn/packet"
)

type PomeloMessagePacket struct{}

func NewPomeloMessagePacket() *PomeloMessagePacket {
	return &PomeloMessagePacket{}
}

// 包编码
func (t *PomeloMessagePacket) Encode(typ packet.Type, data []byte) ([]byte, error) {
	if typ < packet.Handshake || typ > packet.Kick {
		return nil, packet.ErrWrongPomeloPacketType
	}

	if len(data) > MaxPacketSize {
		return nil, ErrPacketSizeExcced
	}

	p := &packet.Packet{Type: typ, Length: len(data)}
	buf := make([]byte, p.Length+HeadLength)
	buf[0] = byte(p.Type)

	copy(buf[1:HeadLength], IntToBytes(p.Length))
	copy(buf[HeadLength:], data)

	return buf, nil
}

// --------------------------------------------------------------------------------------

func (t *PomeloMessagePacket) forward(buf *bytes.Buffer) (int, packet.Type, error) {
	header := buf.Next(HeadLength)
	return ParseHeader(header)
}

// 包解码 Decode decode the network bytes slice to packet.Packet(s)
func (t *PomeloMessagePacket) Decode(data []byte) ([]*packet.Packet, error) {
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
	size, typ, err := t.forward(buf)
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

		size, typ, err = t.forward(buf)
		if err != nil {
			return nil, err
		}
	}

	return packets, nil
}

// ====================================================================

// 消息编码
func (t *PomeloMessagePacket) MessageEncode() {

}

// 消息解码
func (t *PomeloMessagePacket) MessageDecode() {

}
