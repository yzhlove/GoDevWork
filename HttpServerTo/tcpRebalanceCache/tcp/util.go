package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func getSize(reader *bufio.Reader) (size int, err error) {
	var str string
	if str, err = reader.ReadString(' '); err != nil {
		return
	}
	size, err = strconv.Atoi(strings.TrimSpace(str))
	return
}

func setResp(v []byte, ok error, c net.Conn) (err error) {
	str := fmt.Sprintf("%d ", len(v))
	msg := append([]byte(str), v...)
	if ok != nil {
		str = ok.Error()
		msg = []byte(fmt.Sprintf("-%d ", len(str)) + str)
	}
	_, err = c.Write(msg)
	return
}
