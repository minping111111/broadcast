package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore = make(chan bool, 1)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	tcpConn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer tcpConn.Close()
	fmt.Println("connected!")

	go onMessageReceive(tcpConn)

	b := []byte("time\n")
	tcpConn.Write(b)
	<-quitSemaphore

}

func onMessageReceive(conn *net.TCPConn) {

	reader := bufio.NewReader(conn)
	for {

		message, err := reader.ReadString('\n')
		if err != nil {
			break
			quitSemaphore <- true
		}
		fmt.Println(message)
		// var msg string
		// fmt.Println("please input msg...")
		// fmt.Scanln(&msg)
		msg := "客户端1测试。。。"

		time.Sleep(time.Second)
		b := []byte(msg + "\n")
		conn.Write(b)
	}

}
