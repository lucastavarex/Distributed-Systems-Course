package main

import (
	"bufio"
	"fmt"
	"os"

	pipe "github.com/lucastavarex/Distributed-Systems-Course/Pipes"
	signals "github.com/lucastavarex/Distributed-Systems-Course/Signals"
	sockets "github.com/lucastavarex/Distributed-Systems-Course/Sockets"
)

func main() {
	fmt.Printf("[MAIN] Write the program name that you want to run: \n")
	fmt.Printf("       socket_client to run the Socket Client program.\n")
	fmt.Printf("       socket_server to run the Socket Server program.\n")
	fmt.Printf("       signal_rec to run the Signal Receiver program.\n")
	fmt.Printf("       signal_sen to run the Signal Sender program.\n")
	fmt.Printf("       pipe to run the Pipe program.\n\n\n")

	//Collect the input passed by the user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	programName := scanner.Text()

	switch programName {
	case "socket_client":
		sockets.SocketClient()
	case "socket_server":
		sockets.SocketServer()
	case "signal_rec":
		signals.SignalReceiver()
	case "signal_sen":
		signals.SignalSender()
	case "pipe":
		pipe.Pipe()
	default:
		fmt.Println("Incorrect program name")
	}
	return
}
