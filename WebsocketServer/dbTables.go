package WebsocketServer

func (server *WebsocketServer) CreateConnectionsTable() {
	_, err := server.DbConnection.Exec("CREATE TABLE IF NOT EXISTS connections (" +
		"id int PRIMARY KEY AUTO_INCREMENT, " +
		"userId int, " +
		"connectionId CHAR(16), " +
		"createdAt DATETIME(3), " +

		"INDEX (userId), " +
		"INDEX (connectionId)" +
		")")
	if err != nil {
		panic(err)
	}
}

func (server *WebsocketServer) CreateDisconnectsTable() {
	_, err := server.DbConnection.Exec("CREATE TABLE IF NOT EXISTS disconnects (" +
		"id int PRIMARY KEY AUTO_INCREMENT, " +
		"userId int, " +
		"connectionId CHAR(16), " +
		"createdAt DATETIME(3), " +

		"INDEX (userId), " +
		"INDEX (connectionId)" +
		")")
	if err != nil {
		panic(err)
	}
}
