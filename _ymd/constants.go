package _ymd

import "errors"

// Codec constants.
const (
	HeadLength    = 4
	MaxPacketSize = 1 << 24 //16MB
)

// ErrPacketSizeExcced is the error used for encode/decode.
var ErrPacketSizeExcced = errors.New("codec: packet size exceed")
