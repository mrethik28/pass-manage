package storage

import (
	"database/sql"
)

func InitDB() {
	database, _ := sql.Open("sqlite3", "passManage.db")
	initMasterKey, _ := database.Prepare("CREATE TABLE IF NOT EXISTS masterkey (id INTEGER PRIMARY KEY, masterkey TEXT)")
	initPasswords, _ := database.Prepare("CREATE TABLE IF NOT EXISTS passwords (id INTEGER PRIMARY KEY, site TEXT, password TEXT)")
	initMasterKey.Exec()
	initPasswords.Exec()
}
