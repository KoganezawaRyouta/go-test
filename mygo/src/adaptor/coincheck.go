package adaptor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Cticker struct {
	Last      int
	Bid       int
	Ask       int
	High      int
	Low       int
	Volume    string
	Timestamp int
}

type Ctrade struct {
	Id         int
	Amount     string
	Rate       int
	Order_type string
	Created_at string
}

func CoinCheckTicker() Cticker {
	url := "https://coincheck.jp/api/ticker"
	byteArray := get(url)
	var t Cticker
	json.Unmarshal([]byte(string(byteArray)), &t)
	return t
}

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
