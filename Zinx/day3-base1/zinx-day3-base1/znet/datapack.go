package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx-day3-base1/config"
	"zinx-day3-base1/ziface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// ID uint32 4byte + DataLen uint32 4Byte
	return 8
}

func (dp *DataPack) Pack(msg ziface.MessageInterface) ([]byte, error) {

	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dp *DataPack) Unpack(data []byte) (ziface.MessageInterface, error) {

	buf := bytes.NewReader(data)
	msg := &Message{}
	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if size := config.GlobalConfig.MaxPacketSize; size > 0 && msg.DataLen > size {
		return nil, errors.New("too package msg data received")
	}
	return msg, nil
}
