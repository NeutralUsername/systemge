package WebsocketServer

import (
	"HTTPUtilities"
	"ServerUtilities"
	"Utilities"
	"time"
)

type WebsocketServer struct {
	WsListener      *HTTPUtilities.HTTPServer
	ConnectionIdMap map[string]*User
	UserIdMap       map[int]*User

	Base *ServerUtilities.ServerBase
}

func (server *WebsocketServer) ConnectedWebsocketsCount() int {
	server.Base.Mutex.Lock()
	defer server.Base.Mutex.Unlock()

	return len(server.ConnectionIdMap)
}

func (server *WebsocketServer) ConnectedWebsockets() []string {
	server.Base.Mutex.Lock()
	defer server.Base.Mutex.Unlock()

	connections := make([]string, 0, len(server.ConnectionIdMap))
	for connectionId, user := range server.ConnectionIdMap {
		connections = append(connections, "connectionId: \""+connectionId+"\", userId: \""+Utilities.IntToString(user.Id)+"\" since "+user.connections[connectionId].connectedAt.String())
	}
	return connections
}

func (server *WebsocketServer) DisconnectConnections() {
	server.Base.Mutex.Lock()
	defer server.Base.Mutex.Unlock()
	for _, user := range server.UserIdMap {
		for _, connection := range user.connections {
			connection.watchdog.Reset(1 * time.Nanosecond)
		}
	}
}

func CreateWebsocketServer() *WebsocketServer {
	server := &WebsocketServer{
		WsListener:      nil,
		ConnectionIdMap: map[string]*User{},
		UserIdMap:       map[int]*User{},

		Base: ServerUtilities.CreateServerBase(),
	}
	return server
}
