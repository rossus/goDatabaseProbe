package modifier

import (
	"database/sql"
	"github.com/rossus/goDatabaseProbe/common"
	"log"
)

const log_prefix = "Magic: Modifier: "

func Modify(db *sql.DB) {
	statement, err := db.Prepare("INSERT INTO users(name) VALUES ($1)")
	common.CheckError(log_prefix, err)

	//_, err := db.Exec("DELETE FROM users")  // OK
	//_, err := db.Query("DELETE FROM users") // BAD
	result, err := statement.Exec("Neoptolemos")
	common.CheckError(log_prefix, err)

	lastId, err := result.LastInsertId()
	common.WhatDoesItSay(log_prefix, err)

	rowCount, err := result.RowsAffected()
	common.WhatDoesItSay(log_prefix, err)

	log.Println(log_prefix, "ID =", lastId, "; affected =", rowCount)
}