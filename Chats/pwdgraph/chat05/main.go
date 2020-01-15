package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

//短网址算法
//摘要算法

/*
将长网址 md5 生成 32 位签名串,分为 4 段, 每段 8 个字节
对这四段循环处理, 取 8 个字节, 将他看成 16 进制串与 0x3fffffff(30位1) 与操作, 即超过 30 位的忽略处理
这 30 位分成 6 段, 每 5 位的数字作为字母表的索引取得特定字符, 依次进行获得 6 位字符串
总的 md5 串可以获得 4 个 6 位串,取里面的任意一个就可作为这个长 url 的短 url 地址
*/

func main() {
	//transform("www.baidu.com")
	generate("www.baidu.com")
}

func transform(url string) string {

	s := md5.Sum([]byte(url))
	fmt.Printf("%X , len %d \n", s, len(s))

	return ""
}

func generate(url string) string {

	s := md5.Sum([]byte(url))

	fmt.Println("s ==> ", s)

	number := binary.BigEndian.Uint64(s[:])
	fmt.Println("number ===> ", number)

	bytesNumber := make([]byte, 8)

	binary.BigEndian.PutUint64(bytesNumber, number)

	fmt.Println("bytesNumber => ", bytesNumber)

	for _, v := range bytesNumber {
		fmt.Printf("%v %x %.4b \n", v, v, v)
	}

	fmt.Println("=============================")

	t := uint64(0xFF)

	for i := 0; i < 8; i++ {
		fmt.Printf("%v %.4b %x %v\n", number&t, number&t, number&t, byte(number&t))
		number >>= 8
	}

	return ""
}
