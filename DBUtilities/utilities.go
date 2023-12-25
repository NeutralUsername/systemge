package DBUtilities

import (
	"database/sql"
	"fmt"
)

func PingMysqlConnection(sqlCon *sql.DB) bool {
	if err := sqlCon.Ping(); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
