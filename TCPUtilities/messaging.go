package TCPUtilities

import (
	"net"
	"strings"
	"time"
)

func ConstructMessage(messageType string, messageData1 []string, messageData2 []string) []byte {
	return []byte(messageType + TCP_DELIMITER1 + strings.Join(messageData1, TCP_DELIMITER2) + TCP_DELIMITER1 + strings.Join(messageData2, TCP_DELIMITER2))
}

func ParseMessage(message []byte) (string, []string, []string) {
	segments := strings.Split(string(message), TCP_DELIMITER1)
	if len(segments) != 3 {
		return "", nil, nil
	}
	dataSegments1 := strings.Split(segments[1], TCP_DELIMITER2)
	if len(dataSegments1) == 1 && dataSegments1[0] == "" {
		dataSegments1 = nil
	}
	dataSegments2 := strings.Split(segments[2], TCP_DELIMITER2)
	if len(dataSegments2) == 1 && dataSegments2[0] == "" {
		dataSegments2 = nil
	}
	return segments[0], dataSegments1, dataSegments2
}

func MessageConnection(tcpConnection *net.Conn, message []byte) {
	(*tcpConnection).Write(append(message, TCP_ENDOFMESSAGE))
}

func ReadNextMessage(tcpConnection *net.Conn) (message []byte) {
	buffer := make([]byte, BUFFER_SIZE)
	for {
		n, err := (*tcpConnection).Read(buffer)
		if err != nil {
			return nil
		}
		if buffer[n-1] == TCP_ENDOFMESSAGE {
			message = append(message, buffer[:n-1]...)
			break
		} else {
			message = append(message, buffer[:n]...)
		}
	}
	return message
}

func MessageServer(message []byte, address string) error {
	tcpOutgoingConnection, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer tcpOutgoingConnection.Close()

	tcpOutgoingConnection.SetDeadline(time.Now().Add(5 * time.Second))
	MessageConnection(&tcpOutgoingConnection, message)
	return nil
}

func AskServer(message []byte, address string) []byte {
	tcpOutgoingConnection, err := net.Dial("tcp", address)
	if err != nil {
		return nil
	}
	defer tcpOutgoingConnection.Close()

	tcpOutgoingConnection.SetDeadline(time.Now().Add(5 * time.Second))
	MessageConnection(&tcpOutgoingConnection, message)
	replyMessage := ReadNextMessage(&tcpOutgoingConnection)
	return replyMessage
}
