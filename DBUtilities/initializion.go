package DBUtilities

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeConnection(dbUsername string, dbPassword string, dbName string) *sql.DB {
	CreateDatabase(dbUsername, dbPassword, dbName)
	sqlConnection := ConnectToLocalDatabase(dbUsername, dbPassword, dbName)
	sqlConnection.SetMaxOpenConns(100)
	sqlConnection.SetMaxIdleConns(100)
	return sqlConnection
}

func CreateDatabase(username string, password string, dbName string) {
	sqlCon := ConnectToLocalDatabase(username, password, "")
	if err := sqlCon.Ping(); err != nil {
		panic(err)
	}
	_, err := sqlCon.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
	sqlCon.Close()
}

func ConnectToLocalDatabase(username string, password string, dbName string) *sql.DB {
	sqlCon, err := sql.Open("mysql", username+":"+password+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	if !PingMysqlConnection(sqlCon) {
		panic("error in PingMysqlConnection")
	}
	return sqlCon
}
