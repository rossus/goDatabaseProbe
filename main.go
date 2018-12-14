package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/rossus/goDatabaseProbe/common"
	"github.com/rossus/goDatabaseProbe/magic"
)

const log_prefix = "MAIN: "

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=db_probe sslmode=disable")
	common.CheckError(log_prefix, err)
	err = db.Ping()
	common.CheckError(log_prefix, err)

	magic.DoAllMagicPossibleWithThatDBPlease(db)

	defer db.Close()
}
