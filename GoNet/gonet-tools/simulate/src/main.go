package main

import (
	"WorkSpace/GoDevWork/GoNet/gonet-tools/simulate/src/api"
	"WorkSpace/GoDevWork/GoNet/gonet-tools/simulate/src/misc/dh"
	"WorkSpace/GoDevWork/GoNet/gonet-tools/simulate/src/misc/packet"
	"crypto/rc4"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net"
	"time"
)

//////////////////////////////////////////////////////////////////////////
// 封包结构:
//   A	 B	 C	  D
// | 2 | 4 | 2 | data|
// 解释如下:
//	C 和 D 共同组成 payload
// A : package len(2 byte) => len(payload) + seqId(4 byte)
// B : seqId(4 byte)
// C : Code (2 byte)
// D : data (不同的数据大小不一致)
//////////////////////////////////////////////////////////////////////////

var (
	seq         = uint32(0)
	encoder     *rc4.Cipher
	decoder     *rc4.Cipher
	KeyExchange = false
	SALT        = "DH"
)

const (
	DefaultAgentHost = "127.0.0.1:4399"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", DefaultAgentHost)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	S1, M1 := dh.DHExchange()
	S2, M2 := dh.DHExchange()

	seed := api.SeedInfo{ClientSendSeed: int32(M1.Int64()),
		ClientReceiveSeed: int32(M2.Int64())}
	rst := send(conn, api.Code["get_seed_req"], seed)
	tbl, err := api.PacketSeedInfo(rst)
	if err != nil {
		panic(err)
	}
	log.Printf("<<<<<packet seed info>>>>>: %v \n", tbl)

	K1 := dh.DHKey(S1, big.NewInt(int64(tbl.ClientSendSeed)))
	K2 := dh.DHKey(S2, big.NewInt(int64(tbl.ClientReceiveSeed)))

	fmt.Printf(">>>>> K1 %v , K2 %v \n", K1.String(), K2.String())

	encoder, err = rc4.NewCipher([]byte(
		fmt.Sprintf("%v%v", SALT, K1)))
	if err != nil {
		panic(err)
	}
	decoder, err = rc4.NewCipher([]byte(
		fmt.Sprintf("%v%v", SALT, K2)))
	if err != nil {
		panic(err)
	}

	KeyExchange = true

	/*
		user := api.UserLoginInfo{
			LoginWay:          0,
			OpenUid:           "uuid",
			ClientCertificate: "qwertyuiopasdfgh",
			ClientVersion:     1,
			UserLang:          "en",
			AppId:             "com.yzh.love",
			OsVersion:         "android 4.4",
			DeviceName:        "simulate",
			DeviceId:          "device_id",
			DeviceIdType:      1,
			LoginIp:           "127.0.0.1",
		}
	*/

	autoId := api.AutoId{Id: rand.Int31()}
	send(conn, api.Code["heart_beat_req"], autoId)

	autoId = api.AutoId{Id: rand.Int31()}
	send(conn, api.Code["heart_beat_req"], autoId)

	//send_proto(conn, api.Code["user_login_req"], user)

}

func send(conn net.Conn, p int16, info interface{}) (reader *packet.Packet) {
	seq++
	payload := packet.Pack(p, info, nil)
	writer := packet.Writer()
	fmt.Printf("payload length => %v \n", len(payload))
	writer.WriteU16(uint16(len(payload)) + 4)

	//-------------------------------------------------------------------------
	/*
		_reader := packet.Reader(payload)
		code, err := _reader.ReadS16()
		if err != nil {
			panic(err)
		}
		_seed, err := api.PacketSeedInfo(_reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("reader test  code: %v seed: %v \n", code, _seed)
	*/
	//-------------------------------------------------------------------------

	w := packet.Writer()
	w.WriteU32(seq)
	w.WriteRawBytes(payload)
	data := w.Data()

	if KeyExchange {
		encoder.XORKeyStream(data, data)
	}

	writer.WriteRawBytes(data)

	//-------------------------------------------------------------------------
	// 读取封装的包

	newReader := packet.Reader(writer.Data())
	_size, err := newReader.ReadU16()
	if err != nil {
		panic(err)
	}
	fmt.Printf("size = %v \n", _size)
	/*
		_seqId, err := _new_reader.ReadU32()
		if err != nil {
			panic(err)
		}
		fmt.Printf("seqId = %v \n", _seqId)

		_code, err := _new_reader.ReadS16()
		if err != nil {
			panic(err)
		}
		fmt.Printf("code = %v \n", _code)

		_seed_info, err := api.PacketSeedInfo(_new_reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("seed_info = %v \n", _seed_info)
	*/
	//-------------------------------------------------------------------------

	conn.Write(writer.Data())
	log.Printf("send: %#v", writer.Data())
	time.Sleep(time.Second)

	//read
	header := make([]byte, 2)
	io.ReadFull(conn, header)

	size := binary.BigEndian.Uint16(header)
	body := make([]byte, size)
	if _, err := io.ReadFull(conn, body); err != nil {
		panic(err)
	}

	if KeyExchange {
		decoder.XORKeyStream(body, body)
	}

	reader = packet.Reader(body)
	b, err := reader.ReadS16()
	if err != nil {
		panic(err)
	}

	if _, ok := api.RCode[b]; !ok {
		panic(fmt.Sprintf("unknown proto id: %v", b))
	}
	return
}
