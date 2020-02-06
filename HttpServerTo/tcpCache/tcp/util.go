package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func readLen(r *bufio.Reader) (length int, err error) {
	lenStr, err := r.ReadString(' ')
	if err != nil {
		return
	}
	length, err = strconv.Atoi(strings.TrimSpace(lenStr))
	return
}

func sendResp(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		str := fmt.Sprintf("-%d ", len(errString)) + errString
		_, err = conn.Write([]byte(str))
		return err
	}
	valueLength := fmt.Sprintf("%d ", len(value))
	_, err = conn.Write(append([]byte(valueLength), value...))
	return err
}
