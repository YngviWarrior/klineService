package bybitstructs

type GetServerTimestamp struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		TimeNow int64 `json:"serverTime"`
	} `json:"result"`
}

type GetKlinesResponse struct {
	RetCode int64         `json:"ret_code"`
	RetMsg  string        `json:"ret_msg"`
	ExtCode string        `json:"ext_code"`
	ExtInfo string        `json:"ext_info"`
	Result  []interface{} `json:"result"`
}

type CancelOrderResponse struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		AccountId    string `json:"accountId"`
		Symbol       string `json:"symbol"`
		OrderLinkId  string `json:"orderLinkId"`
		OrderId      string `json:"orderId"`
		TransactTime string `json:"transactTime"`
		Price        string `json:"price"`
		OrigQty      string `json:"origQty"`
		ExecutedQty  string `json:"executedQty"`
		Status       string `json:"status"`
		TimeInForce  string `json:"timeInForce"`
		Type         string `json:"type"`
		Side         string `json:"side"`
	} `json:"result"`
}

type OpenOrderResponse struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  []struct {
		AccountId           string `json:"accountId"`
		ExchangeId          string `json:"exchangeId"`
		Symbol              string `json:"symbol"`
		SymbolName          string `json:"symbolName"`
		OrderLinkId         string `json:"orderLinkId"`
		OrderId             string `json:"orderId"`
		Price               string `json:"price"`
		OrigQty             string `json:"origQty"`
		ExecutedQty         string `json:"executedQty"`
		CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
		AvgPrice            string `json:"avgPrice"`
		Status              string `json:"status"`
		TimeInForce         string `json:"timeInForce"`
		Type                string `json:"type"`
		Side                string `json:"side"`
		StopPrice           string `json:"stopPrice"`
		IcebergQty          string `json:"icebergQty"`
		Time                string `json:"time"`
		UpdateTime          string `json:"updateTime"`
		IsWorking           bool   `json:"isWorking"`
	} `json:"result"`
}

type OrderHistoryResponse struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  []struct {
		AccountId           string `json:"accountId"`
		ExchangeId          string `json:"exchangeId"`
		Symbol              string `json:"symbol"`
		SymbolName          string `json:"symbolName"`
		OrderLinkId         string `json:"orderLinkId"`
		OrderId             string `json:"orderId"`
		Price               string `json:"price"`
		OrigQty             string `json:"origQty"`
		ExecutedQty         string `json:"executedQty"`
		CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
		AvgPrice            string `json:"avgPrice"`
		Status              string `json:"status"`
		TimeInForce         string `json:"timeInForce"`
		Type                string `json:"type"`
		Side                string `json:"side"`
		StopPrice           string `json:"stopPrice"`
		IcebergQty          string `json:"icebergQty"`
		Time                string `json:"time"`
		UpdateTime          string `json:"updateTime"`
		IsWorking           bool   `json:"isWorking"`
	} `json:"result"`
}

type AllCoinsResponse struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		Balances []struct {
			Coin     string `json:"coin"`
			CoinId   string `json:"coinId"`
			CoinName string `json:"coinName"`
			Free     string `json:"free"`
			Locked   string `json:"locked"`
			Total    string `json:"total"`
		} `json:"balances"`
	} `json:"result"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int64  `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int64  `json:"rate_limit"`
}

type OrderResponse struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode any    `json:"ext_code"`
	ExtInfo any    `json:"ext_info"`
	Result  struct {
		AccountId    string `json:"accountId"`
		Symbol       string `json:"symbol"`
		SymbolName   string `json:"symbolName"`
		OrderLinkId  string `json:"orderLinkId"`
		OrderId      string `json:"orderId"`
		TransactTime string `json:"transactTime"`
		Price        string `json:"price"`
		OrigQty      string `json:"origQty"`
		ExecutedQty  string `json:"executedQty"`
		Status       string `json:"status"`
		TimeInForce  string `json:"timeInForce"`
		Type         string `json:"type"`
		Side         string `json:"side"`
	} `json:"result"`
}
