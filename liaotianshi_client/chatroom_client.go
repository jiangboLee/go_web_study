package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func MessageSend(conn net.Conn) {
	var input string
	for  {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)

		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Println("client connect failure: " + err.Error())
			break
		}
	}
}

func main() {
	conn, err := net.Dial("tcp","127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go MessageSend(conn)
	//conn.Write([]byte("hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~hello lee~"))

	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("接受到的消息: %s", string(buf[0:length]))
	}
	
	fmt.Println("client end!")
}
