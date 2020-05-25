package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx-day4-base1/config"
	"zinx-day4-base1/ziface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

func (dp *DataPack) Pack(msg ziface.MessageInterface) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMessageID()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (dp *DataPack) Unpack(data []byte) (ziface.MessageInterface, error) {
	buffer := bytes.NewReader(data)
	msg := &Message{}
	if err := binary.Read(buffer, binary.LittleEndian, &msg.Length); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if size := config.GlobalConfig.MaxPackSize; size > 0 && msg.Length >= size {
		return nil, errors.New("overflow package length")
	}
	return msg, nil
}
