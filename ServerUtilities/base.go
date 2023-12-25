package ServerUtilities

import (
	"TCPUtilities"
	"Utilities"
	"database/sql"
	"net"
	"sync"
)

type ServerBase struct {
	TcpListener  *TCPUtilities.Listener
	DbConnection *sql.DB
	Logger       *Utilities.Logger
	Random       *Utilities.Random
	Mutex        *sync.Mutex
}

func CreateServerBase() *ServerBase {
	server := &ServerBase{
		Mutex: &sync.Mutex{},
	}
	return server
}

func (server *ServerBase) InitializeRandom() {
	server.Random = Utilities.CreateRandom()
}

func (server *ServerBase) InitializeLogger(logPath string) {
	server.Logger = Utilities.CreateLogger(logPath)
}

func (server *ServerBase) SetdDbConnection(dbConnection *sql.DB) {
	server.DbConnection = dbConnection
}

func (server *ServerBase) SetTcpListener(port string, handleTcpMessage func(incomingTcpConnection *net.Conn, message []byte)) {
	server.TcpListener = TCPUtilities.CreateListener(port, handleTcpMessage)
}
