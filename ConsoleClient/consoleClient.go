package main

import (
	"System"
	"Utilities"
	"fmt"
)

const (
	EXIT    = "exit"
	START   = "start"
	CLOSE   = "close"
	WSCOUNT = "wscount"
	WSLIST  = "wslist"
	TLS     = "tls"
)

const EXIT_MSG = "exiting console client"
const WELCOME_MSG = "welcome to the console client"
const INPUT_MSG = ">"

func main() {
	println(WELCOME_MSG)
	system := System.CreateSystem(false)
	for {
		print(INPUT_MSG)
		input := readLine()
		switch input {
		case EXIT:
			if system.Status != System.CLOSED {
				system.Close()
			}
			println(EXIT_MSG)
			return
		case START:
			system.Start()
		case CLOSE:
			system.Close()
		case WSCOUNT:
			println("connected websockets : " + Utilities.IntToString(system.WebsocketServer.ConnectedWebsocketsCount()))
		case WSLIST:
			connectedWebsockets := system.WebsocketServer.ConnectedWebsockets()
			for _, connectionId := range connectedWebsockets {
				println(connectionId)
			}
		case TLS:
			if system.Status != System.CLOSED {
				println("servers must be closed before changing TLS")
				break
			}
			system.TLS = !system.TLS
			system = System.CreateSystem(system.TLS)
			println("TLS : " + Utilities.BoolToString(system.TLS))
		default:
			println("unknown command : \"" + input + "\"")
		}
	}
}

func readLine() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	return input
}
