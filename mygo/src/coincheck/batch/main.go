package main

import (
	"coincheck/lib/importer"
	"coincheck/lib/orm"
	"log"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	elapsed := time.Since(time.Now())
	results := importer.CoinCheck()
	var dbmap = orm.InitDb()
	defer dbmap.Db.Close()

	importer.InsertTicker(dbmap, results.Cticker)
	importer.InsertTrade(dbmap, results.Ctrades)

	log.Printf("%s", results.Cticker)
	log.Printf("%s", results.Ctrades)
	log.Printf("%s", elapsed)
}
