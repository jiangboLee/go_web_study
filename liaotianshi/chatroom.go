package main

import (
	"net"
	"fmt"
	"strings"
	"os"
	"log"
)

var onlineConners = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000)
var quitChan = make(chan bool)
var logfile *os.File
var logger *log.Logger
const LOG_DIRECTORY  = "./test.log"
func CheckError(err error)  {
	if err != nil {
		 panic(err)
	}
}

func ProcessInfo(conn net.Conn)  {
	buf := make([]byte, 1024)
	defer func(conn net.Conn) {
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		delete(onlineConners, addr)
		conn.Close()

		for value := range onlineConners {
			fmt.Println("now online conns:" + value)
		}
	}(conn)
	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			message := string(buf[0:numOfBytes])
			messageQueue <- message
		}
	}
}

func ConsumeMessage()  {
	for {
		select {
		case message := <- messageQueue:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}
func doProcessMessage(message string)  {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		//解决信息中也有#的问题
		sendMessage := strings.Join(contents[1:], "#")

		addr = strings.Trim(addr, " ")
		if conn, ok := onlineConners[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failure")
			}
		}
	}
}


func main() {

	logfile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		fmt.Println("log文件开启失败")
		os.Exit(0)
	}
	defer logfile.Close()

	logger = log.New(logfile, "\r\n",log.Ldate|log.Ltime|log.Llongfile)

	logger.Println("我试试logging")

	onlineConners = make(map[string]net.Conn)
	listen_socket, err := net.Listen("tcp","127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close()
	fmt.Println("正zai接听message。。。。。")

	go ConsumeMessage()

	for {
		conn, err := listen_socket.Accept()
		CheckError(err)

		addr := fmt.Sprintf("%s",conn.RemoteAddr())
		onlineConners[addr] = conn
		for i := range onlineConners {
			fmt.Println(i)
		}
		go ProcessInfo(conn)
	}
}
