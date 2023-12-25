package WebsocketServer

import "github.com/gorilla/websocket"

func (WebsocketConnection *Connection) message(message []byte) {
	WebsocketConnection.mutex.Lock()
	defer WebsocketConnection.mutex.Unlock()

	WebsocketConnection.connection.WriteMessage(websocket.TextMessage, message)
}

func (server *WebsocketServer) messageConnectionByConnectionId(connectionId string, message []byte) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if user := server.ConnectionIdMap[connectionId]; user != nil {
		if connection := user.connections[connectionId]; connection != nil {
			go connection.message(message)
		}
	}
}

func (server *WebsocketServer) messageUserByConnectionId(connectionId string, message []byte) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if user := server.ConnectionIdMap[connectionId]; user != nil {
		for _, connection := range user.connections {
			go connection.message(message)
		}
	}
}

func (server *WebsocketServer) messageUserIds(userIds []int, message []byte) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for _, userId := range userIds {
		if user := server.UserIdMap[userId]; user != nil {
			for _, connection := range user.connections {
				go connection.message(message)
			}
		}
	}
}
