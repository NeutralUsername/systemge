package WebsocketServer

import (
	"MessageTypes"
	"Ports"
	"TCPUtilities"
	"Utilities"
	"sync"
)

func (server *WebsocketServer) assembleLoginData(loginUserId int) (loginData []string) {
	var wg sync.WaitGroup
	reuqestCount := 1
	makeRequest := func(messageType string, port string) {
		defer wg.Done()
		if responseMessageType, messageData1, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(messageType, nil, []string{Utilities.IntToString(loginUserId)}), port)); responseMessageType == MessageTypes.TRUE {
			switch messageType {
			case MessageTypes.GETUSERPRIVATE:
				loginData[0] = messageData1[0]
			}
		}
	}
	loginData = make([]string, reuqestCount)
	wg.Add(reuqestCount)
	go makeRequest(MessageTypes.GETUSERPRIVATE, Ports.USERSERVER_TCP_LISTENER)
	wg.Wait()
	return loginData
}
