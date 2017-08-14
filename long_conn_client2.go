package main

import (
	"bufio"
	"fmt"
	"net"
	// "time"
)

var quitSemaphore = make(chan bool, 1)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "10.74.36.12:9999")

	tcpConn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer tcpConn.Close()
	fmt.Println("connected!")

	go onMessageReceive(tcpConn)

	// b := []byte("time\n")
	// tcpConn.Write(b)
	<-quitSemaphore

}

func onMessageReceive(conn *net.TCPConn) {

	reader := bufio.NewReader(conn)
	go func() {
		var msg string
		for {

			// fmt.Scanln(&msg)
			fmt.Print("我说：")
			fmt.Scan(&msg)
			b := []byte(msg + "\n")
			conn.Write(b)
		}
	}()
	for {

		message, err := reader.ReadString('\n')
		if err != nil {
			break
			quitSemaphore <- true
		}
		fmt.Println("                                               " + "她说：" + message)
		// fmt.Println(message)
		// var msg string
		// fmt.Println("please input msg...")
		// fmt.Scanln(&msg)

		// msg := "客户端1测试。。。"

		// time.Sleep(time.Second)
		// b := []byte(msg + "\n")
		// conn.Write(b)
	}

}
