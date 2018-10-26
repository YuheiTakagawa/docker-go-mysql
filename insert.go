package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

func main() {
	dbmap := initDB()
	defer dbmap.Db.Close()

	// ダミートランザクションを作成
	t := time.Now()
	const layout2 = "2006-01-02 15:04:05.00"
	timestamp := t.Format(layout2)
	dummy := Transaction{ID: 0, TxHash: "ebi", BlockID: 1, Input: "input", Output: "output", Amount: 10, Timestamp: timestamp, Sign: "sign", Pubkey: "pubkey"}

	// トランザクションの記録先テーブルのindexを決定する関数
	index := chooseDB(dummy)
	//トランザクションとindexを渡してinsertする関数
	insert(dummy, index)

	// var transactions []Transaction
	// _, err := dbmap.Select(&transactions, "select * from transactions order by id")
	// checkErr(err, "Select failed")
	// log.Println("All rows:")
	// for x, p := range transactions {
	// 	log.Printf("    %d: %v\n", x, p)
	// }
}

func initDB() *gorp.DbMap {
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/db")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbmap.AddTableWithName(Transaction{}, "transactions").SetKeys(true, "id")
	dbmap.AddTableWithName(Table1{}, "table1").SetKeys(true, "id")
	dbmap.AddTableWithName(Table2{}, "table2").SetKeys(true, "id")
	dbmap.AddTableWithName(Table3{}, "table3").SetKeys(true, "id")
	dbmap.AddTableWithName(Table4{}, "table4").SetKeys(true, "id")
	dbmap.AddTableWithName(RingBuffer{}, "ring_buffer").SetKeys(true, "id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func chooseDB(tx Transaction) int {
	dbmap := initDB()
	defer dbmap.Db.Close()
	fmt.Println(tx.Input)
	var index RingBuffer
	err := dbmap.SelectOne(&index, "select * from ring_buffer where user_address=?", tx.Input)

	// RingBufferテーブルにデータがなかったら新規作成
	if err != nil {
		new_address := &RingBuffer{1, tx.Input, 1}
		dbmap.Insert(new_address)
		return new_address.TableAddress
	}

	// 読み込んだアドレスに+1してindexとして返す
	table_address := index.TableAddress
	update_index := &RingBuffer{index.ID, index.UserAddress, index.TableAddress}
	update_index.TableAddress = table_address + 1
	if update_index.TableAddress == 5 {
		update_index.TableAddress = 1
	}
	dbmap.Update(update_index)
	return update_index.TableAddress
}

func insert(tx Transaction, index int) {
	dbmap := initDB()
	if index == 1 {
		var insert_tx Table1
		insert_tx.ID = tx.ID
		insert_tx.TxHash = tx.TxHash
		insert_tx.BlockID = tx.BlockID
		insert_tx.Input = tx.Input
		insert_tx.Output = tx.Output
		insert_tx.Amount = tx.Amount
		insert_tx.Timestamp = tx.Timestamp
		insert_tx.Sign = tx.Sign
		insert_tx.Pubkey = tx.Pubkey
		fmt.Println(insert_tx)
		err := dbmap.Insert(&insert_tx)
		checkErr(err, "Insert failed")
	}
	if index == 2 {
		var insert_tx Table2
		insert_tx.ID = tx.ID
		insert_tx.TxHash = tx.TxHash
		insert_tx.BlockID = tx.BlockID
		insert_tx.Input = tx.Input
		insert_tx.Output = tx.Output
		insert_tx.Amount = tx.Amount
		insert_tx.Timestamp = tx.Timestamp
		insert_tx.Sign = tx.Sign
		insert_tx.Pubkey = tx.Pubkey
		fmt.Println(insert_tx)
		err := dbmap.Insert(&insert_tx)
		checkErr(err, "Insert failed")
	}
	if index == 3 {
		var insert_tx Table3
		insert_tx.ID = tx.ID
		insert_tx.TxHash = tx.TxHash
		insert_tx.BlockID = tx.BlockID
		insert_tx.Input = tx.Input
		insert_tx.Output = tx.Output
		insert_tx.Amount = tx.Amount
		insert_tx.Timestamp = tx.Timestamp
		insert_tx.Sign = tx.Sign
		insert_tx.Pubkey = tx.Pubkey
		fmt.Println(insert_tx)
		err := dbmap.Insert(&insert_tx)
		checkErr(err, "Insert failed")
	}
	if index == 4 {
		var insert_tx Table4
		insert_tx.ID = tx.ID
		insert_tx.TxHash = tx.TxHash
		insert_tx.BlockID = tx.BlockID
		insert_tx.Input = tx.Input
		insert_tx.Output = tx.Output
		insert_tx.Amount = tx.Amount
		insert_tx.Timestamp = tx.Timestamp
		insert_tx.Sign = tx.Sign
		insert_tx.Pubkey = tx.Pubkey
		fmt.Println(insert_tx)
		err := dbmap.Insert(&insert_tx)
		checkErr(err, "Insert failed")
	}
}

type Transaction struct {
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

type Table1 struct {
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

type Table2 struct {
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

type Table3 struct {
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

type Table4 struct {
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

type RingBuffer struct {
	ID           int    `db:"id"`
	UserAddress  string `db:"user_address"`
	TableAddress int    `db:"table_address"`
}
