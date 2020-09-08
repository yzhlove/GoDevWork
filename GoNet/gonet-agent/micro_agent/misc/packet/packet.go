package packet

import "errors"

const (
	PACKET_LIMIT = 65535
)

var PosIndexRange = errors.New("pos index range")
var ReadIndexRange = errors.New("read index range")

type Packet struct {
	pos  int
	data []byte
}

func (p *Packet) Data() []byte {
	return p.data
}

func (p *Packet) Length() int {
	return len(p.data)
}

func (p *Packet) ReadBool() (bool, error) {
	if b, err := p.ReadByte(); err != nil {
		return false, err
	} else {
		if b != byte(1) {
			return false, nil
		}
		return true, nil
	}
}

func (p *Packet) ReadByte() (ret byte, err error) {
	if p.pos >= len(p.data) {
		err = errors.New("read byte failed")
		return
	}
	ret = p.data[p.pos]
	p.pos++
	return
}

func (p *Packet) ReadBytes() (ret []byte, err error) {
	if p.pos+2 > len(p.data) {
		err = errors.New("read bytes header failed")
		return
	}
	size, err := p.ReadU16()
	if err != nil {
		return
	}

	if p.pos+int(size) > len(p.data) {
		err = errors.New("read bytes data failed")
		return
	}

	ret = p.data[p.pos : p.pos+int(size)]
	p.pos += int(size)
	return
}

func (p *Packet) ReadString() (ret string, err error) {
	if p.pos+2 >
}

func (p *Packet) ReadU16() (ret uint16, err error) {
	if p.pos+2 > len(p.data) {
		err = errors.New("read uint16 failed")
		return
	}
	buf := p.data[p.pos : p.pos+2]
	ret = uint16(buf[0])<<8 | uint16(buf[1])
	p.pos += 2
	return
}
