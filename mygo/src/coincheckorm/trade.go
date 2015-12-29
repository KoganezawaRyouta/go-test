package coincheckorm

type Trade struct {
	ID        int    `db:"id, primarykey, autoincrement"`
	TradeID   int    `db:"trade_id"`
	Amount    string `db:"amount"`
	Rate      int    `db:"rate"`
	OrderType string `db:"order_type"`
	CreatedAt string `db:"created_at"`
}
