package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/config"
	"zinx/ziface"
)

type Package struct{}

func NewPack() *Package {
	return &Package{}
}

func (Package) HeadSize() uint32 {
	//ID 4 + data length 4
	return 8
}

func (Package) Pack(msg ziface.MsgImp) (data []byte, err error) {
	sb := bytes.NewBuffer([]byte{})
	if err = binary.Write(sb, binary.LittleEndian, msg.GetSize()); err != nil {
		return
	}
	if err = binary.Write(sb, binary.LittleEndian, msg.GetID()); err != nil {
		return
	}
	if err = binary.Write(sb, binary.LittleEndian, msg.GetData()); err != nil {
		return
	}
	return sb.Bytes(), nil
}

func (Package) Unpack(data []byte) (ziface.MsgImp, error) {
	sb := bytes.NewReader(data)
	msg := &Msg{}
	if err := binary.Read(sb, binary.LittleEndian, &msg.Size); err != nil {
		return nil, err
	}
	if err := binary.Read(sb, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if size := config.GlobalConfig.MaxPackageSize; size <= msg.Size {
		return nil, errors.New("overflow max package size")
	}
	return msg, nil
}
