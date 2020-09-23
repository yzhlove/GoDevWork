package handler

import (
	"crypto/rc4"
	"errors"
	"fmt"
	"math/big"
	"micro_agent/misc/dh"
	"micro_agent/misc/packet"
	"micro_agent/sess"
)

//密钥交换
//加密方式建立: DH + RC4
//完整的加密过程包括: RSA + DH + RC4
//1. RSA用于鉴定服务器的真伪
//2. DH用于在不安全的信道上协商安全的key
//3. RC4用于流加密
func GetSeedReq(s *sess.Session, reader *packet.Packet) []byte {
	tbl, err := PacketSeedInfo(reader)
	if err != nil {
		return failed(s, errors.New("packet err:"+err.Error()))
	}

	fmt.Printf("read seed info => %v \n", tbl)

	//KEY1
	X1, E1 := dh.DHExchange()
	KEY1 := dh.DHKey(X1, big.NewInt(int64(tbl.ClientSendSeed)))

	//KEY2
	X2, E2 := dh.DHExchange()
	KEY2 := dh.DHKey(X2, big.NewInt(int64(tbl.ClientReceiveSeed)))

	fmt.Printf(">>>>>> E1:%v E2:%v KEY1:%v KEY2:%v \n", E1.Int64(), E2.Int64(), KEY1.String(), KEY2.String())

	ret := SeedInfo{ClientSendSeed: int32(E1.Int64()), ClientReceiveSeed: int32(E2.Int64())}
	//服务器加密种子是客户端解密种子
	encoder, err := rc4.NewCipher([]byte(fmt.Sprintf("%v%v", Salt, KEY2)))
	if err != nil {
		return failed(s, errors.New("rc4 encoder err:"+err.Error()))
	}

	decoder, err := rc4.NewCipher([]byte(fmt.Sprintf("%v%v", Salt, KEY1)))
	if err != nil {
		return failed(s, errors.New("rc4 decoder err:"+err.Error()))

	}

	s.Encoder = encoder
	s.Decoder = decoder
	s.Flag |= sess.SESS_KEYEXCG
	return succeed(s, ret)
}
