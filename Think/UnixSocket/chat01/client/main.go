package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	var file string = "test.sock"

	conn, err := net.Dial("unix", file)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(conn)

	for {

		fmt.Print("input text:")
		input.Scan()

		conn.Write(input.Bytes())
		conn.Write([]byte("\n"))

		data, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("read data:", data)
	}

}
