package main

import (
	"coincheck/importer"
	"log"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	elapsed := time.Since(time.Now())
	results := importer.CoinCheck()

	importer.InsertTicker(results.Cticker)
	importer.InsertTrade(results.Ctrades)

	log.Printf("%s", results.Cticker)
	log.Printf("%s", results.Ctrades)
	log.Printf("%s", elapsed)
}
