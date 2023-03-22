package binancestructs

type OrderParams struct {
	Symbol string
	Side   string
	Type   string
	// timeInForce 	ENUM 	NO
	Quantity      float64 // 1º coin of the parity
	QuoteOrderQty float64 // 2º coin of the parity
	// price 	DECIMAL 	NO
	// newClientOrderId string  	NO 	A unique id among open orders. Automatically generated if not sent.
	// strategyId 	INT 	NO
	// strategyType 	INT 	NO 	The value cannot be less than 1000000.
	// stopPrice 	DECIMAL 	NO 	Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	// trailingDelta 	LONG 	NO 	Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. For more details on SPOT implementation on trailing stops, please refer to Trailing Stop FAQ
	// icebergQty 	DECIMAL 	NO 	Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
	// newOrderRespType 	ENUM 	NO 	Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.

}
