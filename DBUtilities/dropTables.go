package DBUtilities

import "database/sql"

func DropAllTables(db *sql.DB) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			panic(err)
		}
		_, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0")
		if err != nil {
			panic(err)
		}
		DropTable(db, table)
		_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
		if err != nil {
			panic(err)
		}
	}
}

func DropTable(db *sql.DB, tableName string) {
	_, err := db.Exec("DROP TABLE IF EXISTS " + tableName)
	if err != nil {
		panic(err)
	}
}
