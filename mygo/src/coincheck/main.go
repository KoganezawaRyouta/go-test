package main

import (
	"adaptor"
	"coincheckorm"
	"database/sql"
	"log"
	"math/rand"
	"time"

	"gopkg.in/gorp.v1"
)

type CoinCheckResult struct {
	Cticker adaptor.Cticker
	Ctrades []adaptor.Ctrade
}

var (
	db, _ = sql.Open("sqlite3", "/coincheckdb/coincheck.db")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
)

func main() {

	defer dbmap.Db.Close()
	rand.Seed(time.Now().UnixNano())
	elapsed := time.Since(time.Now())
	results := coinCheck()

	insertTicker(results.Cticker)
	insertTrade(results.Ctrades)

	log.Printf("%s", results.Cticker)
	log.Printf("%s", results.Ctrades)
	log.Printf("%s", elapsed)
}

func coinCheck() (chResult CoinCheckResult) {
	chTicker := make(chan adaptor.Cticker)
	chTrades := make(chan []adaptor.Ctrade)
	timeout := time.After(500 * time.Millisecond)

	go func() { chTicker <- adaptor.CoinCheckTicker() }()
	go func() { chTrades <- adaptor.CoinCheckTrades() }()

	for i := 0; i < 2; i++ {
		select {
		case result := <-chTicker:
			chResult.Cticker = result
		case result := <-chTrades:
			chResult.Ctrades = result
		case <-timeout:
			log.Printf("%s", "timed out")
			return
		}
	}
	return
}

func insertTicker(cticker adaptor.Cticker) {
	ticker := newTicker(cticker)
	err1 := dbmap.Insert(&ticker)
	checkErr(err1, "Insert failed ticker")
}

func insertTrade(ctrades []adaptor.Ctrade) {
	for _, ctrade := range ctrades {
		trades := newTrade(ctrade)
		err2 := dbmap.Insert(&trades)
		checkErr(err2, "Insert failed trades")
	}
}

func newTicker(chTicker adaptor.Cticker) coincheckorm.Ticker {
	return coincheckorm.Ticker{
		Last:      chTicker.Last,
		Bid:       chTicker.Bid,
		Ask:       chTicker.Ask,
		High:      chTicker.High,
		Low:       chTicker.Low,
		Volume:    chTicker.Volume,
		Timestamp: chTicker.Timestamp,
	}
}

func newTrade(chTrade adaptor.Ctrade) coincheckorm.Trade {
	return coincheckorm.Trade{
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
