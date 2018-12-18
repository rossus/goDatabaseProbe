package retriever

import (
	"database/sql"
	"github.com/rossus/goDatabaseProbe/common"
	"log"
)

func SingleRowQuery(db *sql.DB) {

	var name string

	err := db.QueryRow("select name from users where id = $1", 3).Scan(&name)
	common.CheckError(log_prefix, err)
	log.Println(log_prefix, name)


	statement, err := db.Prepare("select name from users where id = $1")
	common.CheckError(log_prefix, err)
	defer statement.Close()

	var anotherName string
	err = statement.QueryRow(4).Scan(&anotherName)
	common.CheckError(log_prefix, err)
	log.Println(log_prefix, anotherName)
}