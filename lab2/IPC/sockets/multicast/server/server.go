// Very basic socket server
// https://golangr.com/

package main

import (
	"bufio"
	"fmt"
	"net"
)
var sockets []net.Conn
func broadcast( socket net.Conn) {
	for {
		message, _ := bufio.NewReader(socket).ReadString('\n')
		for _, conn := range sockets {
			if (conn ==socket){continue}
			fmt.Println("Message Received:", string(message))
			fmt.Fprintf(conn, message+"\n")
		}
	}
}

func main() {
	fmt.Println("Start server...")
	//sockets := []net.Conn{}
	serverSocket, _ := net.Listen("tcp", ":8000")
	for {
		// accept connection
		conn, _ := serverSocket.Accept()
		sockets = append(sockets, conn)
		go broadcast(conn)
	}
}
