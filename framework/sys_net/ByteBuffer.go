package sys_net

import (
	"bytes"
	"encoding/binary"
	"phoenix-tudou/framework/log4"
)

type ByteBuffer struct {
	mBuf *bytes.Buffer
}

/*
	ByteBuffer

added by yh @ 2023-04-24
*/
func NewBytebuffer() *ByteBuffer {
	var bb = &ByteBuffer{}
	bb.mBuf = &bytes.Buffer{}
	//buffer := bytes.NewBufferString("Hello, world!")
	return bb
}

// 正序 binary.BigEndian 反序binary.LittleEndian
func (t *ByteBuffer) SetData(bytes []byte) {
	binary.Write(t.mBuf, binary.BigEndian, &bytes)
}

func (t *ByteBuffer) Clear() {
	t.mBuf = new(bytes.Buffer)
}
func (t *ByteBuffer) Close() {
	t.mBuf = nil
}

// #region ======byte Read and write=======================================
func (t *ByteBuffer) WriteByte(v byte) {
	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadByte() byte {
	var tmp byte
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

// 添加二进制数据
func (t *ByteBuffer) WriteBytes(bytes []byte) {
	t.WriteInt32(int32(len(bytes)))
	binary.Write(t.mBuf, binary.BigEndian, &bytes)
}

// 添加二进制数据
func (t *ByteBuffer) ReadBytes() []byte {
	l := t.ReadInt32()
	if l <= 0 {
		return []byte{}
	}
	if int(l) > t.mBuf.Len() {
		log4.Error("ByteBuffer->ReadBytes l=%d", l)
		return []byte{}
	}
	bytes := make([]byte, l)
	binary.Read(t.mBuf, binary.BigEndian, &bytes)
	return bytes
}

//#endregion =============================================================

//#region ======String Read and write=====================================

// 最大字节 65535
func (t *ByteBuffer) WriteString(v string) {
	bytes := []byte(v)
	nLen := len(bytes)
	t.WriteUInt16(uint16(nLen))
	binary.Write(t.mBuf, binary.BigEndian, bytes)
}
func (t *ByteBuffer) ReadString() string {
	nLen := t.ReadUInt16()
	bytes := make([]byte, nLen)
	binary.Read(t.mBuf, binary.BigEndian, &bytes)
	var tmp string
	tmp = string(bytes[:])
	return tmp
}

//#endregion =============================================================

//#region ====Reading and writing double-precision floating-point numbers===========

// -------float------------------
func (t *ByteBuffer) WriteFloat32(v float32) {

	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadFloat32() float32 {
	var tmp float32
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}
func (t *ByteBuffer) WriteFloat64(v float64) {

	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadFloat64() float64 {
	var tmp float64
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

//#endregion=======双精度浮点数读写====================================================

//#region ====int read and write==========================================

// -------int16------------------
func (t *ByteBuffer) WriteInt16(v int16) {
	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadInt16() int16 {
	var tmp int16
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

// -------int32------------------
func (t *ByteBuffer) WriteInt32(v int32) {
	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadInt32() int32 {
	var tmp int32
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

// -------int64------------------
func (t *ByteBuffer) WriteInt64(v int64) {
	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadInt64() int64 {
	var tmp int64
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

//#endregion ================================================================

//#region ====Unsigned int read and write==========================================

// -------int16------------------
func (t *ByteBuffer) WriteUInt16(v uint16) {
	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadUInt16() uint16 {
	var tmp uint16
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

// -------int32------------------
func (t *ByteBuffer) WriteUInt32(v uint32) {
	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadUInt32() uint32 {
	var tmp uint32
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

// -------int64------------------
func (t *ByteBuffer) WriteUInt64(v uint64) {
	binary.Write(t.mBuf, binary.BigEndian, &v)
}
func (t *ByteBuffer) ReadUInt64() uint64 {
	var tmp uint64
	binary.Read(t.mBuf, binary.BigEndian, &tmp)
	return tmp
}

//#endregion ================================================================

func (t *ByteBuffer) ToBytes() []byte {
	//var tmp []byte
	//copy(t.m_buf.Bytes(),tmp)
	return t.mBuf.Bytes()
	//return t.m_buf.Flush();
}

/*
func (t *ByteBuffer) IsEmpty() bool {
	return reflect.DeepEqual(t,ByteBuffer{})
}*/
