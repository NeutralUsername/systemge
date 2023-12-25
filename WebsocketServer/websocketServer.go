package WebsocketServer

import (
	"HTTPUtilities"
	"TCPUtilities"
	"Utilities"
	"database/sql"
	"sync"
	"time"
)

type WebsocketServer struct {
	WsListener      *HTTPUtilities.HTTPServer
	ConnectionIdMap map[string]*User
	UserIdMap       map[int]*User
	mutex           *sync.Mutex

	TcpListener  *TCPUtilities.Listener
	DbConnection *sql.DB
	Logger       *Utilities.Logger
	Random       *Utilities.Random
}

func (server *WebsocketServer) ConnectedWebsocketsCount() int {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	return len(server.ConnectionIdMap)
}

func (server *WebsocketServer) ConnectedWebsockets() []string {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	connections := make([]string, 0, len(server.ConnectionIdMap))
	for connectionId, user := range server.ConnectionIdMap {
		connections = append(connections, "connectionId: \""+connectionId+"\", userId: \""+Utilities.IntToString(user.Id)+"\" since "+user.connections[connectionId].connectedAt.String())
	}
	return connections
}

func (server *WebsocketServer) DisconnectConnections() {
	server.mutex.Lock()
	defer server.mutex.Unlock()
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
		mutex:           &sync.Mutex{},
	}
	return server
}
