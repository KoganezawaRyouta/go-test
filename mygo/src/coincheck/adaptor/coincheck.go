package adaptor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Cticker Structure of ticker by api result
type Cticker struct {
	Last      int
	Bid       int
	Ask       int
	High      int
	Low       int
	Volume    string
	Timestamp int
}

// Ctrade Structure of trades by api result
type Ctrade struct {
	Id         int
	Amount     string
	Rate       int
	Order_type string
	Created_at string
}

// CoinCheckTicker get ticker from coincheck.jp
func CoinCheckTicker() Cticker {
	url := "https://coincheck.jp/api/ticker"
	byteArray := get(url)
	var t Cticker
	json.Unmarshal([]byte(string(byteArray)), &t)
	return t
}

// CoinCheckTrades get trades from coincheck.jp
func CoinCheckTrades() []Ctrade {
	url := "https://coincheck.jp/api/trades"
	byteArray := get(url)
	var t []Ctrade
	json.Unmarshal([]byte(string(byteArray)), &t)
	return t
}

func get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray
}
