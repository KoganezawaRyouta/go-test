package coincheckorm

type Ticker struct {
	ID        int    `db:"id, primarykey, autoincrement"`
	Last      int    `db:"last"`
	Bid       int    `db:"bid"`
	Ask       int    `db:"ask"`
	High      int    `db:"high"`
	Low       int    `db:"low"`
	Volume    string `db:"volume"`
	Timestamp int    `db:"timestamp"`
}
