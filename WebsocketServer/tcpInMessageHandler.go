package WebsocketServer

import (
	"TCPUtilities"
	"Utilities"
	"net"
)

func (server *WebsocketServer) HandleTcpMessage(incomingTcpConnection *net.Conn, message []byte) {
	messageType, messageData1, messageData2 := TCPUtilities.ParseMessage(message)
	switch messageType {
	default:
		server.messageUserIds(Utilities.StringsToInts(messageData2), TCPUtilities.ConstructMessage(messageType, messageData1, nil))
	}
}
