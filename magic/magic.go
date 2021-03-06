package magic

import (
	"database/sql"
	"github.com/rossus/goDatabaseProbe/magic/modifier"
	"github.com/rossus/goDatabaseProbe/magic/retriever"
)

func DoAllMagicPossibleWithThatDBPlease(db *sql.DB) {

	//I. Retrieving Result Sets
	//
	//Based on http://go-database-sql.org/retrieving.html
	//
	//There are several idiomatic operations to retrieve results from the datastore.
	//1. Execute a query that returns rows:
	retriever.FetchDataFrom(db)
	//2. Prepare a statement for repeated use, execute it multiple times, and destroy it:
	retriever.PrepareQuery(db)
	//3. Execute a statement in a once-off fashion, without preparing it for repeated use.
	//4. Execute a query that returns a single row. There is a shortcut for this special case:
	retriever.SingleRowQuery(db)

	//II. Modifying Data and Using Transactions
	//
	//Based on http://go-database-sql.org/modifying.html
	//
	//1. Statements that Modify Data:
	modifier.Modify(db)
	//2. Working with Transactions.
}