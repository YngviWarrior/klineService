package kline

type Kline struct {
	Asset            uint64  `json:"asset"`
	AssetSymbol      string  `json:"asset_symbol"`
	AssetQuote       uint64  `json:"asset_quote"`
	AssetQuoteSymbol string  `json:"asset_quote_symbol"`
	Exchange         uint64  `json:"exchange"`
	Mts              uint64  `json:"mts"`
	Open             float64 `json:"open"`
	Close            float64 `json:"close"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Volume           float64 `json:"volume"`
	TestNet          bool    `json:"testnet"`
	Roc              float64 `json:"roc"`
}
