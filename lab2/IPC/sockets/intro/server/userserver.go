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
	fmt.Println("Client :",mess+"\n")
}

}

func outGoingHandler(socket net.Conn)  {
 for{
	 fmt.Println("Message : ")
	serveMessage,_:=bufio.NewReader(os.Stdin).ReadString('\n')
	 fmt.Fprint(socket,serveMessage+"\n")

 }
}


func main(){
	l, _ := net.Listen("tcp", "127.0.0.1:8000")
      fmt.Println("connecting ...")
	sock, _ := l.Accept()
	fmt.Println("connected")
	var wg =sync.WaitGroup{}
	wg.Add(1)

	go incommingHandler(sock)
	go outGoingHandler(sock)
	wg.Wait()



}
