package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

type ReceivedTransaction struct {
	ID        int    `db:"id"`
	TxHash    string `db:"TxHash"`
	BlockID   int    `db:"BlockID"`
	Input     string `db:"Input"`
	Output    string `db:"Output"`
	Amount    int    `db:"Amount"`
	Timestamp string `db:"Timestamp"`
	Sign      string `db:"Sign"`
	Pubkey    string `db:"Pubkey"`
}

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()
	tx1 := &ReceivedTransaction{0, "hoge", 1, "input", "output", 10, "time", "sign", "pubkey"}
	err := dbmap.Insert(tx1)
	checkErr(err, "Insert failed")

	var transactions []ReceivedTransaction
	_, err = dbmap.Select(&transactions, "select * from txs order by id")
	checkErr(err, "Select failed")
	log.Println("All rows:")
	for x, p := range transactions {
		log.Printf("    %d: %v\n", x, p)
	}
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/db")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(ReceivedTransaction{}, "txs").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
