package WebsocketServer

import "time"

func (server *WebsocketServer) insertConnection(userId int, connectionId string, createdAt time.Time) {
	_, err := server.Base.DbConnection.Exec("INSERT INTO connections ("+
		"userId, connectionId, createdAt"+
		") VALUES (?, ?, ?)",
		userId, connectionId, createdAt)
	if err != nil {
		server.Base.Logger.Log(err.Error())
	}
}

func (server *WebsocketServer) insertDisconnect(userId int, connectionId string, createdAt time.Time) {
	_, err := server.Base.DbConnection.Exec("INSERT INTO disconnects ("+
		"userId, connectionId, createdAt"+
		") VALUES (?, ?, ?)",
		userId, connectionId, createdAt)
	if err != nil {
		server.Base.Logger.Log(err.Error())
	}
}
