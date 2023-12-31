package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	//TODO In a continuous loop, read a message from the server and display it.
	for {
		//read a message
		msg, _ := reader.ReadString('\n')
		//display
		fmt.Println(msg)
	}
}

func write(conn *net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("enter Text: ")
		msg, _ := reader.ReadString('\n')
		if msg == "/quit\n" {
			break
		} else {
			fmt.Fprintln(*conn, msg)
		}
	}

}

func main() {

	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr)

	go read(&conn)
	write(&conn)

	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
