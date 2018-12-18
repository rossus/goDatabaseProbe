package retriever

import (
	"database/sql"
	"github.com/rossus/goDatabaseProbe/common"
	"log"
)

func PrepareQuery(db *sql.DB) {

	var (
		id int
		name string
	)

	statement, err := db.Prepare("select id, name from users where id = $1")
	common.CheckError(log_prefix, err)
	defer statement.Close()

	rows, err := statement.Query(2)
	common.CheckError(log_prefix, err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name)
		common.CheckError(log_prefix, err)
		log.Println(log_prefix, id, name)
	}

	err = rows.Err()
	common.CheckError(log_prefix, err)
}