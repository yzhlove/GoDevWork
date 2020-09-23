package main

import (
	"crypto/rc4"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"time"
)

//服务端的底数和素数是共享的
var (
	PRIME, _ = big.NewInt(0).SetString("0x7FFFFFC3", 0)        //素数
	BASE     = big.NewInt(3)                                   //底数
	E        = rand.New(rand.NewSource(time.Now().UnixNano())) //随机数
	MAXINT64 = big.NewInt(math.MaxInt64)                       //随机数范围限定
)

var (
	slat = "dh"
)

func main() {

	//client 客户端行为

	//客户端先生成两组密钥（公钥和私钥）
	clientX1, clientE1 := change()
	clientX2, clientE2 := change()

	//客户端把公钥发送给服务器（clientE1 and clientE2)

	//server 服务端行为

	//服务端也生成两组密钥
	serverX1, serverE1 := change()
	serverX2, serverE2 := change()

	//服务端拿到客户端的公钥算出KEY值
	serverKEY1 := key(serverX1, clientE1)
	serverKEY2 := key(serverX2, clientE2)

	log.Printf("[server] key1:%v key2:%v \n", serverKEY1.Int64(), serverKEY2.Int64())

	//接着，服务端根据算出来的key基于rc4做加密解密
	// * 服务端的加密种子是客户端的解密种子
	server_encoder, err := rc4.NewCipher([]byte(fmt.Sprintf("%v%v", slat, serverKEY1.String())))
	if err != nil {
		panic(err)
	}

	server_decoder, err := rc4.NewCipher([]byte(fmt.Sprintf("%v%v", slat, serverKEY2.String())))
	if err != nil {
		panic(err)
	}

	//client 客户端行为

	//客户端拿到服务器的公钥算出key值
	clientKEY1 := key(clientX1, serverE1)
	clientKEY2 := key(clientX2, serverE2)

	log.Printf("[client] key1:%v key2:%v \n", clientKEY1.Int64(), clientKEY2.Int64())

	if serverKEY1.Cmp(clientKEY1) != 0 {
		panic("server and client key1 no comp")
	}

	if serverKEY2.Cmp(clientKEY2) != 0 {
		panic("server and client key2 no comp")
	}

	//客户端根据算出来的key值基于rc4做加密解密
	client_decoder, err := rc4.NewCipher([]byte(fmt.Sprintf("%v%v", slat, clientKEY1)))
	if err != nil {
		panic(err)
	}

	client_encoder, err := rc4.NewCipher([]byte(fmt.Sprintf("%v%v", slat, clientKEY2)))
	if err != nil {
		panic(err)
	}

	//接着就是客户端和服务器相互发送数据的加密与解密过程

	// 客户端发送数据

	//客户端消息体
	client_message := []byte("hello world server !!!")

	//客户端加密数据
	client_encoder.XORKeyStream(client_message, client_message)
	log.Printf("client encoder message:%v \n", string(client_message))

	//服务端解密客户端加密的数据
	server_decoder.XORKeyStream(client_message, client_message)
	log.Printf("server decoder message:%v \n", string(client_message))

	//服务端发送数据

	//服务端消息
	server_message := []byte("what are you doing client ???")

	//服务端加密数据
	server_encoder.XORKeyStream(server_message, server_message)
	log.Printf("server encoder message:%v \n", string(server_message))

	//客户端解密 服务端加密过的数据
	client_decoder.XORKeyStream(server_message, server_message)
	log.Printf("client decoder message:%v \n", string(server_message))

}

// 测试算法用
func mock() {

	// A
	X1, E1 := change()

	// B
	X2, E2 := change()

	// A ------

	KEY1 := key(X1, E2)

	// B ------
	KEY2 := key(X2, E1)

	// result
	fmt.Printf("key1:%v key2:%v comp:%v \n", KEY1.Int64(), KEY2.Int64(), KEY1.Cmp(KEY2))

}

// change 返回私钥和公钥
func change() (*big.Int, *big.Int) {
	_secret := big.NewInt(0).Rand(E, MAXINT64)
	_public := big.NewInt(0).Exp(BASE, _secret, PRIME)
	return _secret, _public
}

func key(secret, public *big.Int) *big.Int {
	k := big.NewInt(0).Exp(public, secret, PRIME)
	return k
}
