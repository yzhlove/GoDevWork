package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func readSize(r *bufio.Reader) (size int, err error) {
	var sizeStr string
	if sizeStr, err = r.ReadString(' '); err != nil {
		return
	}
	size, err = strconv.Atoi(strings.TrimSpace(sizeStr))
	return
}

func sendResponse(value []byte, status error, conn net.Conn) (err error) {
	sizeStr := fmt.Sprintf("%d ", len(value))
	message := append([]byte(sizeStr), value...)
	if status != nil {
		str := status.Error()
		message = []byte(fmt.Sprintf("-%d ", len(str)) + str)
	}
	_, err = conn.Write(message)
	return
}
