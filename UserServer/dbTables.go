package UserServer

func (server *UserServer) CreateUsersTable() {
	_, err := server.Base.DbConnection.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id int PRIMARY KEY AUTO_INCREMENT, " +
		"username VARCHAR(50) COLLATE utf8_bin UNIQUE, " +
		"password CHAR(64), " +
		"email VARCHAR(200) UNIQUE NULL, " +
		"secretQuestion TEXT NULL, " +
		"secretAnswer CHAR(64) NULL, " +
		"userPower INT, " +
		"createdAt DATETIME(3), " +
		"registeredAt DATETIME(3), " +
		"confirmedAt DATETIME(3)," +

		"showCommunicator BOOLEAN, " +
		"communicatorContent INT," +
		"showTutorial BOOLEAN, " +
		"uiTheme INT, " +

		"INDEX (username)" +
		")")
	if err != nil {
		panic(err)
	}
}
