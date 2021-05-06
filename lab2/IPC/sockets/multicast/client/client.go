package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)
func incommingHandler(socket net.Conn)  {
	for{
		mess,_:=bufio.NewReader(socket).ReadString('\n')
		fmt.Println("Server:",mess)
	}

}

func outGoingHandler(socket net.Conn)  {
	for{
		fmt.Println("Message : ")
		serveMessage,_:=bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprint(socket,serveMessage+"\n")
	}
}
func main() {
	// connect to server
	socket, _ := net.Dial("tcp", "127.0.0.1:8000")
	for {
		var wg =sync.WaitGroup{}
		wg.Add(1)
		go incommingHandler(socket)
		go outGoingHandler(socket)
		wg.Wait()
	}
}
