package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var ConnMap = make(map[string]*net.TCPConn)

func main() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		fmt.Println("a client connected..." + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("dis connected..." + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

	for {

		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		// fmt.Println(message)
		// msg := time.Now().String() + "\n"
		time.Sleep(time.Second)
		b := []byte(message)
		for _, v := range ConnMap {
			v.Write(b)
		}
		// conn.Write(b)
	}
}
