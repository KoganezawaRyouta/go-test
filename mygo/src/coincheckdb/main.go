package main

import (
	"coincheckorm"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

func main() {
	os.Remove("./cointcheck.db")
	db, err := sql.Open("sqlite3", "./cointcheck.db")
	if err != nil {
		panic(err.Error())
	}

	dbmap := gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	ticker := dbmap.AddTableWithName(coincheckorm.Ticker{}, "ticker").SetKeys(true, "Id")
	ticker.ColMap("ID").Rename("id")
	ticker.ColMap("Last").Rename("last")
	ticker.ColMap("Bid").Rename("bid")
	ticker.ColMap("Ask").Rename("ask")
	ticker.ColMap("Low").Rename("low")
	ticker.ColMap("Volume").Rename("volume")
	ticker.ColMap("Timestamp").Rename("timestamp")

	trade := dbmap.AddTableWithName(coincheckorm.Trade{}, "trade").SetKeys(true, "Id")
	trade.ColMap("ID").Rename("id")
	trade.ColMap("TradeID").Rename("trade_id")
	trade.ColMap("Amount").Rename("amount")
	trade.ColMap("Rate").Rename("rate")
	trade.ColMap("OrderType").Rename("order_type")
	trade.ColMap("CreatedAt").Rename("created_at")

	dbmap.DropTables()
	err = dbmap.CreateTables()
	if err != nil {
		panic(err.Error())
	}
}
