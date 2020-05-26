package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx-day6-base1/config"
	"zinx-day6-base1/ziface"
)

type Package struct{}

func (p *Package) MaxHead() uint32 {
	return 8
}

func (p *Package) Pack(msg ziface.MsgInterface) ([]byte, error) {
	sb := bytes.NewBuffer([]byte{})
	if err := binary.Write(sb, binary.LittleEndian, msg.GetDataSize()); err != nil {
		return nil, err
	}
	if err := binary.Write(sb, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	if err := binary.Write(sb, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return sb.Bytes(), nil
}

func (p *Package) Unpack(data []byte) (ziface.MsgInterface, error) {
	sb := bytes.NewReader(data)
	msg := &Msg{}
	if err := binary.Read(sb, binary.LittleEndian, &msg.Size); err != nil {
		return nil, err
	}
	if err := binary.Read(sb, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if size := config.GlobalConfig.MaxPackSize; size > 0 && size < msg.Size {
		return nil, errors.New("overflow package size ")
	}
	return msg, nil
}
