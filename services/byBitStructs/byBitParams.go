package bybitstructs

type GetKlinesParams struct {
	Test      bool
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	Limit     int64  `json:"limit"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

type CancelOrderParams struct {
	Test        bool
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type OpenOrderParams struct {
	Test    bool
	Symbol  string `json:"symbol"`
	OrderId string `json:"orderId"`
	Limit   int64  `json:"limit"`
}

type OrderHistoryParams struct {
	Test      bool
	Symbol    string `json:"symbol"`
	OrderId   string `json:"orderId"`
	Limit     int64  `json:"limit"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

type OrderParams struct {
	Test          bool
	OrderId       string `json:"orderId"`
	OrderLinkId   string `json:"orderLinkId"`
	Symbol        string `json:"symbol"`
	CreateTime    string `json:"createTime"`
	OrderPrice    string `json:"orderPrice"`
	OrderQty      string `json:"orderQty"`
	OrderType     string `json:"orderType"`
	Side          string `json:"side"`
	Status        string `json:"status"`
	TimeInForce   string `json:"timeInForce"`
	AccountId     string `json:"accountId"`
	OrderCategory int64  `json:"orderCategory"`
	TriggerPrice  string `json:"triggerPrice"`
}
