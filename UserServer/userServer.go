package UserServer

import (
	"TCPUtilities"
	"Utilities"
	"database/sql"
)

type UserServer struct {
	TcpListener  *TCPUtilities.Listener
	DbConnection *sql.DB
	Logger       *Utilities.Logger
	Random       *Utilities.Random
}

func CreateUserServer() *UserServer {
	server := &UserServer{}
	return server
}
