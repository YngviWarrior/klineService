package binancestructs

type TraderFee struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

type OrderResponse struct {
	Code                int64  `json:"code"`
	Msg                 string `json:"msg"`
	Symbol              string `json:"symbol"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StrategyId          int64  `json:"strategyId"`
	StrategyType        int64  `json:"strategyType"`
	Fills               []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int64  `json:"tradeId"`
	}
}

type AllCoinsResponse struct {
	Coin             string `json:"coin"`
	DepositAllEnable bool   `json:"depositAllEnable"`
	Free             string `json:"free"`
	Freeze           string `json:"freeze"`
	Ipoable          string `json:"ipoable"`
	Ipoing           string `json:"ipoing"`
	IsLegalMoney     bool   `json:"isLegalMoney"`
	Locked           string `json:"locked"`
	Name             string `json:"name"`
	NetworkList      []struct {
		AddressRegex            string `json:"addressRegex"`
		Coin                    string `json:"coin"`
		DepositDesc             string `json:"depositDesc"`
		DepositEnable           bool   `json:"depositEnable"`
		IsDefault               bool   `json:"isDefault"`
		MemoRegex               string `json:"memoRegex"`
		MinConfirm              int64  `json:"minConfirm"`
		Name                    string `json:"name"`
		Network                 string `json:"network"`
		ResetAddressStatus      bool   `json:"resetAddressStatus"`
		SpecialTips             string `json:"specialTips"`
		UnLockConfirm           int64  `json:"unLockConfirm"`
		WithdrawDesc            string `json:"withdrawDesc"`
		WithdrawEnable          bool   `json:"withdrawEnable"`
		WithdrawFee             string `json:"withdrawFee"`
		WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
		WithdrawMax             string `json:"withdrawMax"`
		WithdrawMin             string `json:"withdrawMin"`
		SameAddress             bool   `json:"sameAddress"`
		EstimatedArrivalTime    int64  `json:"estimatedArrivalTime"`
		Busy                    bool   `json:"busy"`
	} `json:"networkList"`
	Storage           string `json:"storage"`
	Trading           bool   `json:"trading"`
	WithdrawAllEnable bool   `json:"withdrawAllEnable"`
	Withdrawing       string `json:"withdrawing"`
}
