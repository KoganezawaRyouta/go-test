package importer

import (
	"coincheck/adaptor"
	"coincheck/orm"
	"database/sql"
	"log"
	"time"

	"gopkg.in/gorp.v1"
)

// CoinCheckResult obtained from the coincheck.jp
type CoinCheckResult struct {
	Cticker adaptor.Cticker
	Ctrades []adaptor.Ctrade
}

var (
	db, _ = sql.Open("sqlite3", "/coincheckdb/coincheck.db")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
)

// CoinCheck it obtains the information of the trades and ticker from coincheck.jp,
// and register to DB
func CoinCheck() (chResult CoinCheckResult) {
	chTicker := make(chan adaptor.Cticker)
	chTrades := make(chan []adaptor.Ctrade)

	go func() { chTicker <- adaptor.CoinCheckTicker() }()
	go func() { chTrades <- adaptor.CoinCheckTrades() }()

	for i := 0; i < 2; i++ {
		select {
		case result := <-chTicker:
			chResult.Cticker = result
		case result := <-chTrades:
			chResult.Ctrades = result
		case <-time.After(500 * time.Millisecond):
			log.Printf("%s", "timed out")
			return
		}
	}
	return
}

// InsertTicker db insert to orm.Ticker
func InsertTicker(cticker adaptor.Cticker) {
	defer dbmap.Db.Close()
	ticker := newTicker(cticker)
	err := dbmap.Insert(&ticker)
	checkErr(err, "Insert failed ticker")
}

// InsertTrade db insert to orm.Trade
func InsertTrade(ctrades []adaptor.Ctrade) {
	defer dbmap.Db.Close()
	for _, ctrade := range ctrades {
		trades := newTrade(ctrade)
		err := dbmap.Insert(&trades)
		checkErr(err, "Insert failed trades")
	}
}

func newTicker(chTicker adaptor.Cticker) orm.Ticker {
	return orm.Ticker{
		Last:      chTicker.Last,
		Bid:       chTicker.Bid,
		Ask:       chTicker.Ask,
		High:      chTicker.High,
		Low:       chTicker.Low,
		Volume:    chTicker.Volume,
		Timestamp: chTicker.Timestamp,
	}
}

func newTrade(chTrade adaptor.Ctrade) orm.Trade {
	return orm.Trade{
		TradeID:   chTrade.Id,
		Amount:    chTrade.Amount,
		Rate:      chTrade.Rate,
		OrderType: chTrade.Order_type,
		CreatedAt: chTrade.Created_at,
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
