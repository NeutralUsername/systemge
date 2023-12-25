package WebsocketServer

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	Id          int
	connections map[string]*Connection
}

func CreateUser(id int) *User {
	return &User{
		Id:          id,
		connections: map[string]*Connection{},
	}
}

type Connection struct {
	id          string
	user        *User
	connection  *websocket.Conn
	mutex       *sync.Mutex
	watchdog    *time.Timer
	connectedAt time.Time
}

func (server *WebsocketServer) CreateConnection(user *User, connectionId string, conn *websocket.Conn) *Connection {
	connection := &Connection{
		id:          connectionId,
		user:        user,
		connection:  conn,
		mutex:       &sync.Mutex{},
		watchdog:    nil,
		connectedAt: time.Now(),
	}
	connection.watchdog = time.AfterFunc(WEBSOCKET_WATCHDOG_TIMEOUT, func() { server.removeConnection(connection) })
	return connection
}

func (server *WebsocketServer) removeConnection(connection *Connection) {
	connection.watchdog.Stop()
	connection.connection.Close()

	server.Base.Mutex.Lock()
	defer server.Base.Mutex.Unlock()

	user := connection.user
	connection.user = nil

	delete(user.connections, connection.id)
	delete(server.ConnectionIdMap, connection.id)
	if len(user.connections) == 0 {
		delete(server.UserIdMap, user.Id)
	}

	go server.insertDisconnect(user.Id, connection.id, time.Now())
}

func (server *WebsocketServer) addConnection(userId int, wsConnection *websocket.Conn) *Connection {
	server.Base.Mutex.Lock()
	defer server.Base.Mutex.Unlock()

	connectionId := server.Base.Random.GenerateRandomString(CONNECTION_ID_LENGTH, VALID_CONNECTION_ID_CHARS)
	for server.ConnectionIdMap[connectionId] != nil {
		connectionId = server.Base.Random.GenerateRandomString(CONNECTION_ID_LENGTH, VALID_CONNECTION_ID_CHARS)
	}

	go server.insertConnection(userId, connectionId, time.Now())

	if user := server.UserIdMap[userId]; user != nil {
		connection := server.CreateConnection(user, connectionId, wsConnection)
		server.ConnectionIdMap[connectionId] = user
		user.connections[connectionId] = connection
		return connection
	} else {
		user := CreateUser(userId)
		connection := server.CreateConnection(user, connectionId, wsConnection)
		server.UserIdMap[userId] = user
		server.ConnectionIdMap[connectionId] = user
		user.connections[connectionId] = connection
		return connection
	}
}

func (server *WebsocketServer) changeUser(connection *Connection, userId int) {
	server.Base.Mutex.Lock()
	defer server.Base.Mutex.Unlock()

	oldUser := server.ConnectionIdMap[connection.id]
	delete(server.ConnectionIdMap, connection.id)
	delete(oldUser.connections, connection.id)
	if len(oldUser.connections) == 0 {
		delete(server.UserIdMap, oldUser.Id)
	}

	go server.insertDisconnect(oldUser.Id, connection.id, time.Now())
	go server.insertConnection(userId, connection.id, time.Now())

	if existingNewUser := server.UserIdMap[userId]; existingNewUser != nil {
		connection.user = existingNewUser
		server.ConnectionIdMap[connection.id] = existingNewUser
		existingNewUser.connections[connection.id] = connection
	} else {
		newUser := CreateUser(userId)
		connection.user = newUser
		server.UserIdMap[userId] = newUser
		server.ConnectionIdMap[connection.id] = newUser
		newUser.connections[connection.id] = connection
	}
}
