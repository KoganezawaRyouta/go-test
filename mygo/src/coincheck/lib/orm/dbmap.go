package orm

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

var DatabaseFile = "./cointcheck.db"

func InitDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		panic(err.Error())
	}

	dbmap := gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	ticker := dbmap.AddTableWithName(Ticker{}, "ticker").SetKeys(true, "ID")
	ticker.ColMap("ID").Rename("id")
	ticker.ColMap("Last").Rename("last")
	ticker.ColMap("Bid").Rename("bid")
	ticker.ColMap("Ask").Rename("ask")
	ticker.ColMap("Low").Rename("low")
	ticker.ColMap("Volume").Rename("volume")
	ticker.ColMap("Timestamp").Rename("timestamp")

	trade := dbmap.AddTableWithName(Trade{}, "trade").SetKeys(true, "ID")
	trade.ColMap("ID").Rename("id")
	trade.ColMap("TradeID").Rename("trade_id")
	trade.ColMap("Amount").Rename("amount")
	trade.ColMap("Rate").Rename("rate")
	trade.ColMap("OrderType").Rename("order_type")
	trade.ColMap("CreatedAt").Rename("created_at")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		panic(err.Error())
	}

	return &dbmap
}
